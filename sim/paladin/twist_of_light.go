package paladin

import (
	"github.com/wowsims/classic/sim/core"
)

const (
	echoOfCommand       int32 = 1311703
	echoOfRighteousness int32 = 1311704
)

type sealEchoSource struct {
	id    int32
	spell *core.Spell
}

type sealEcho struct {
	aura  *core.Aura
	spell *core.Spell
}

func (echo *sealEcho) consume(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
	if !result.Landed() || !spell.ProcMask.Matches(core.ProcMaskMeleeWhiteHit) {
		return
	}
	proc := echo.spell
	echo.spell = nil
	aura.Deactivate(sim)
	if proc != nil {
		proc.Cast(sim, result.Target)
	}
}

// Twist of Light leaves one separately charged Echo per replaced seal.
// Client aura options for both Echoes specify one charge, no duration, and
// a white-melee proc mask; Judgements do not consume them.
func (paladin *Paladin) registerTwistOfLight() {
	if !paladin.Talents.TwistOfLight {
		return
	}

	command := &sealEcho{}
	command.aura = paladin.RegisterAura(core.Aura{
		Label:     "Echo of Command",
		ActionID:  core.ActionID{SpellID: 1311703},
		Duration:  core.NeverExpires,
		MaxStacks: 1,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			aura.SetStacks(sim, 1)
		},
		OnSpellHitDealt: command.consume,
	})
	righteousness := &sealEcho{}
	righteousness.aura = paladin.RegisterAura(core.Aura{
		Label:     "Echo of Righteousness",
		ActionID:  core.ActionID{SpellID: 1311704},
		Duration:  core.NeverExpires,
		MaxStacks: 1,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			aura.SetStacks(sim, 1)
		},
		OnSpellHitDealt: righteousness.consume,
	})
	paladin.sealEchoes = map[int32]*sealEcho{
		echoOfCommand:       command,
		echoOfRighteousness: righteousness,
	}
}

func (paladin *Paladin) registerSealProc(seal *core.Aura, proc *core.Spell, echoID int32) {
	if paladin.sealProcs == nil {
		paladin.sealProcs = make(map[*core.Aura]sealEchoSource)
	}
	paladin.sealProcs[seal] = sealEchoSource{id: echoID, spell: proc}
}

func (paladin *Paladin) bankSealEcho(sim *core.Simulation, newSeal *core.Aura) {
	if paladin.sealEchoes == nil || newSeal == paladin.currentSeal || !paladin.currentSeal.IsActive() {
		return
	}
	source, ok := paladin.sealProcs[paladin.currentSeal]
	if !ok {
		return
	}
	echo := paladin.sealEchoes[source.id]
	if echo == nil {
		return
	}
	echo.spell = source.spell
	echo.aura.Activate(sim)
}
