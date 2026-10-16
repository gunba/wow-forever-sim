package planner

import (
	"fmt"
	"sync"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
	"github.com/wowsims/classic/sim/core/stats"
	"github.com/wowsims/classic/sim/warrior"
	"github.com/wowsims/classic/sim/warrior/ranks"
	googleProto "google.golang.org/protobuf/proto"
)

// Scenario is one benchmark fight: a character, a pull, a partner and the console's intent.
type Scenario struct {
	Name       string
	Player     *proto.Player
	Enemies    int
	EnemyLevel int32
	Elite      bool
	HPMult     float64 // enemy health relative to the level's expected creature health
	DmgMult    float64 // enemy damage relative to the level's expected creature damage
	Partner    Partner
	Intent     Intent
	HomeStance int
	HomeWeapon int
	TTKNoise   float64
	Cap        float64  // seconds; the fight ends here if enemies are left
	Front      bool     // the Warrior attacks from the front (enemy parry and block)
	Manual     []string // abilities kept manual besides the defaults
	TTKUnknown float64  // seconds after an enemy first takes damage before the console knows its time to die
}

// Outcome is one iteration's result.
type Outcome struct {
	KillTime, Useful, Taken, LossTime, Threat, LossFocus float64
	Swaps, Decisions                                     int
	ChooseNanos                                          int64
	Casts                                                map[string]int
	DamageBy                                             map[string]float64
}

var setupMu sync.Mutex

func init() {
	warrior.CustomRotationHook = hook
	warrior.ExecuteRangeHook = func(sim *core.Simulation, target *core.Unit) bool {
		h := harnessOf(target)
		if h == nil {
			return sim.IsExecutePhase20()
		}
		i := h.index(target)
		return i >= 0 && !h.Dead[i] && (h.ForceExecute || h.HealthFrac(i) <= 0.2)
	}
}

var (
	harnessMu sync.Mutex
	harnesses = map[*core.Unit]*Harness{}
)

func harnessOf(u *core.Unit) *Harness {
	harnessMu.Lock()
	defer harnessMu.Unlock()
	return harnesses[u]
}

// Bench is a scenario built once and run many times with one chooser.
type Bench struct {
	sim     *core.Simulation
	driver  *Driver
	harness *Harness
	Model   *Model
}

// NewBench builds the scenario's simulation. choose makes the chooser for the driver.
func NewBench(sc Scenario, model *Model, choose func(d *Driver) Chooser) (*Bench, error) {
	player := googleProto.Clone(sc.Player).(*proto.Player)
	player.InFrontOfTarget = sc.Front
	player.Rotation = &proto.APLRotation{Type: proto.APLRotation_TypeAPL, PriorityList: []*proto.APLListItem{
		{Action: &proto.APLAction{Action: &proto.APLAction_CustomRotation{CustomRotation: &proto.APLActionCustomRotation{}}}}}}
	lvl := min(max(sc.EnemyLevel, 1), int32(len(expectedCreatures)-1))
	ec := expectedCreatures[lvl]
	hp := ec.health * max(sc.HPMult, 0.01)
	dmg := ec.autoAttackDPS * 2 / 1.17 * max(sc.DmgMult, 0)
	var targets []*proto.Target
	for i := 0; i < max(sc.Enemies, 1); i++ {
		targets = append(targets, &proto.Target{Level: sc.EnemyLevel, MobType: proto.MobType_MobTypeHumanoid,
			Stats: stats.Stats{stats.Armor: ec.armor, stats.Health: hp}.ToFloatArray(), MinBaseDamage: dmg, DamageSpread: 0.3333, SwingSpeed: 2})
	}
	enc := &proto.Encounter{Duration: max(sc.Cap, 30), Targets: targets}
	raid := core.SinglePlayerRaidProto(player, &proto.PartyBuffs{}, &proto.RaidBuffs{}, &proto.Debuffs{})
	raid.Tanks = []*proto.UnitReference{{Type: proto.UnitReference_Player, Index: 0}}
	rsr := &proto.RaidSimRequest{Raid: raid, Encounter: enc, SimOptions: &proto.SimOptions{Iterations: 1, RandomSeed: 1, Ruleset: proto.Ruleset_RulesetForever}}
	hps := make([]float64, len(targets))
	for i := range hps {
		hps[i] = hp
	}
	b := &Bench{Model: model}
	group := 1
	if sc.Partner.Role != "" {
		group = 2
	}
	cfg := DriveConfig{Intent: sc.Intent, HomeStance: sc.HomeStance, TTKNoise: sc.TTKNoise, Elite: sc.Elite,
		PriorDPS: PriorGroupDPS(float64(sc.Player.Level), group), TTKUnknown: sc.TTKUnknown}
	setupMu.Lock()
	warrior.SetupHook = func(w *warrior.Warrior) {
		b.harness = newHarness(w, sc.Partner, hps)
		cfg.Harness = b.harness
		b.driver = newDriver(w, cfg)
		b.driver.model = model
	}
	func() {
		defer func() {
			warrior.SetupHook = nil
			setupMu.Unlock()
		}()
		b.sim = core.NewSim(rsr, simsignals.CreateSignals())
	}()
	if b.driver == nil {
		return nil, fmt.Errorf("the scenario has no Warrior")
	}
	harnessMu.Lock()
	for _, u := range b.sim.Encounter.TargetUnits {
		harnesses[u] = b.harness
	}
	harnessMu.Unlock()
	for _, name := range sc.Manual {
		b.driver.manual[ranks.Ability(name)] = true
	}
	b.driver.chooser = choose(b.driver)
	return b, nil
}

// Run runs iteration i (its seed offset) and returns the outcome.
func (b *Bench) Run(i int64) Outcome {
	d := b.driver
	d.Decisions, d.ChooseNanos, d.Casts = 0, 0, nil
	b.sim.RunIteration(i)
	h := b.harness
	kt := h.KillTime
	if kt == 0 {
		kt = b.sim.CurrentTime.Seconds()
	}
	threat := 0.0
	for _, v := range h.My {
		threat += v
	}
	out := Outcome{KillTime: kt, Useful: h.UsefulDamage(), Taken: h.Taken, LossTime: h.LossTime, Swaps: h.Swaps, Threat: threat, LossFocus: h.LossFocus,
		Decisions: d.Decisions, ChooseNanos: d.ChooseNanos, Casts: map[string]int{}}
	for k, v := range d.Casts {
		out.Casts[string(k)] = v
	}
	out.DamageBy = map[string]float64{}
	for k, v := range h.DamageBy {
		name := fmt.Sprintf("spell %d", k)
		switch {
		case k == -1:
			name = "whiteMH"
		case k == -2:
			name = "whiteOH"
		default:
			if a, r := ranks.ByID(k); r != nil {
				name = string(a)
			}
		}
		out.DamageBy[name] += v
	}
	return out
}

// Driver is the bench's console emulator.
func (b *Bench) Driver() *Driver { return b.driver }

// Release drops the bench's registrations.
func (b *Bench) Release() {
	harnessMu.Lock()
	for _, u := range b.sim.Encounter.TargetUnits {
		delete(harnesses, u)
	}
	harnessMu.Unlock()
	driversMu.Lock()
	delete(drivers, b.driver.w)
	driversMu.Unlock()
}

// ModelFor builds the planner's model from a player proto, the way the live planner does from a character message.
func ModelFor(player *proto.Player) (*Model, error) { return buildFromPlayer(player) }

func cloneStance(p *proto.Player, stance int) *proto.Player {
	q := googleProto.Clone(p).(*proto.Player)
	st := proto.WarriorStance(stance)
	if st == 0 {
		st = proto.WarriorStance_WarriorStanceBattle
	}
	q.Spec = &proto.Player_TankWarrior{TankWarrior: &proto.TankWarrior{Options: &proto.TankWarrior_Options{Stance: st}}}
	return q
}

// PriorGroupDPS is the console's damage-rate guess before a fight shows one: a Warrior's damage per second at the
// level (0.015 L^2 + 1.4 L, fitted to the bench's priority-list fights at levels 10-60) for each player, with each
// member after the first counted at 0.7 (tanks and healers deal less). The time to die blends it with the observed
// rate, so the estimate exists from the pull (no unknown period) and gives way to the measured rate within seconds.
func PriorGroupDPS(level float64, groupSize int) float64 {
	if level <= 0 {
		level = 60
	}
	return (0.015*level*level + 1.4*level) * (1 + 0.7*float64(max(groupSize, 1)-1))
}
