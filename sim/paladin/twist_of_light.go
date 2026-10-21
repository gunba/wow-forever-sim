package paladin

import (
	"github.com/wowsims/classic/sim/core"
)

const (
	echoOfCommand       int32 = 1311703
	echoOfRighteousness int32 = 1311704
	echoOfFury          int32 = 1311701
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
	fury := &sealEcho{}
	fury.aura = paladin.RegisterAura(core.Aura{
		Label: "Echo of Fury", ActionID: core.ActionID{SpellID: echoOfFury},
		Duration: core.NeverExpires, MaxStacks: 1,
		OnGain:          func(aura *core.Aura, sim *core.Simulation) { aura.SetStacks(sim, 1) },
		OnSpellHitDealt: fury.consume,
	})
	paladin.sealEchoes = map[int32]*sealEcho{
		echoOfCommand:       command,
		echoOfRighteousness: righteousness,
		echoOfFury:          fury,
	}
}

func (paladin *Paladin) registerSealProc(seal *core.Aura, proc *core.Spell, echoID int32) {
	if paladin.sealProcs == nil {
		paladin.sealProcs = make(map[*core.Aura]sealEchoSource)
	}
	paladin.sealProcs[seal] = sealEchoSource{id: echoID, spell: proc}
}

// On Forever only one Echo is up at a time: gaining one drops any other (wowsims/forever #218, seen in game), and a
// seal and its own Echo are never up together, so casting a seal drops its Echo and a new rank of the same seal
// leaves none (#124).
func (paladin *Paladin) bankSealEcho(sim *core.Simulation, newSeal *core.Aura) {
	if paladin.sealEchoes == nil {
		return
	}
	forever := paladin.Env.IsForever()
	if newSeal != paladin.currentSeal && paladin.currentSeal.IsActive() {
		if source, ok := paladin.sealProcs[paladin.currentSeal]; ok {
			if echo := paladin.sealEchoes[source.id]; echo != nil {
				if forever {
					for _, other := range paladin.sealEchoes {
						if other != echo {
							other.drop(sim)
						}
					}
				}
				echo.spell = source.spell
				echo.aura.Activate(sim)
			}
		}
	}
	if forever {
		if own, ok := paladin.sealProcs[newSeal]; ok {
			if echo := paladin.sealEchoes[own.id]; echo != nil {
				echo.drop(sim)
			}
		}
	}
}

func (echo *sealEcho) drop(sim *core.Simulation) {
	echo.spell = nil
	if echo.aura.IsActive() {
		echo.aura.Deactivate(sim)
	}
}
