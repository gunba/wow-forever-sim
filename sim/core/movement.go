package core

import (
	"container/heap"
	"math"
	"time"

	"github.com/wowsims/classic/sim/core/proto"
)

type MoveModifier struct {
	ActionId *ActionID
	Modifier float64
}

type MoveHeap []MoveModifier

func (h MoveHeap) Len() int           { return len(h) }
func (h MoveHeap) Less(i, j int) bool { return h[i].Modifier > h[j].Modifier }
func (h MoveHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MoveHeap) Push(x any) {
	*h = append(*h, x.(MoveModifier))
}

func (h *MoveHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *MoveHeap) activeModifier() float64 {
	heap := *h
	if len(heap) < 1 {
		panic("Movement Heaps should never be missing their base case!!!")
	}
	n := heap[0]
	return n.Modifier
}

func (h *MoveHeap) Find(actionId *ActionID) int {
	for i, mod := range *h {
		if mod.ActionId == actionId {
			return i
		}
	}
	return -1
}

func (move *MovementHandler) addMoveSpeedModifier(moveHeap *MoveHeap, moveMod MoveModifier) {
	heap.Push(moveHeap, moveMod)
	move.updateMoveSpeed()
}

func (move *MovementHandler) removeMoveSpeedModifier(moveHeap *MoveHeap, actionID *ActionID) {
	index := moveHeap.Find(actionID)
	if index == -1 {
		return
	}
	heap.Remove(moveHeap, index)
	move.updateMoveSpeed()
}

func (move *MovementHandler) updateMoveSpeed() {
	move.MoveSpeed = move.baseSpeed * move.getActveModifier(move.moveSpeedBonuses) * (1 - move.getActveModifier(move.moveSpeedPenalties))
}

func (move *MovementHandler) getActveModifier(moveHeap *MoveHeap) float64 {
	return moveHeap.activeModifier()
}

type MovementHandler struct {
	Moving    bool
	MoveSpeed float64

	unit        *Unit
	rootCount   int
	moveRange   float64
	lastUpdated time.Duration
	activeSpeed float64
	moveAction  *PendingAction

	baseSpeed          float64
	moveAura           *Aura
	moveSpell          *Spell
	moveSpeedBonuses   *MoveHeap
	moveSpeedPenalties *MoveHeap
}

func (unit *Unit) initMovement() {
	unit.MovementHandler = &MovementHandler{
		unit: unit,
		moveSpeedBonuses: &MoveHeap{
			MoveModifier{
				Modifier: 1,
			},
		},
		moveSpeedPenalties: &MoveHeap{
			MoveModifier{
				Modifier: 0,
			},
		},
		baseSpeed: 7.0,
	}
	unit.MovementHandler.updateMoveSpeed()

	unit.MovementHandler.moveAura = unit.GetOrRegisterAura(Aura{
		Label:     "Movement",
		ActionID:  ActionID{OtherID: proto.OtherAction_OtherActionMove},
		Duration:  NeverExpires,
		MaxStacks: 30,

		OnGain: func(aura *Aura, sim *Simulation) {
			if unit.IsChanneling(sim) {
				unit.ChanneledDot.Cancel(sim)
			}
			unit.AutoAttacks.CancelAutoSwing(sim)
			unit.MovementHandler.Moving = true
		},
		OnExpire: func(aura *Aura, sim *Simulation) {
			move := unit.MovementHandler
			if move.moveAction != nil {
				move.moveAction.Cancel(sim)
				move.moveAction = nil
			}
			move.Moving, move.activeSpeed = false, 0
			if unit.IsEnabled() && !unit.PseudoStats.Stunned {
				unit.AutoAttacks.EnableAutoSwing(sim)
				// Simulate the delay from starting attack.
				unit.AutoAttacks.DelayMeleeBy(sim, time.Millisecond*50)
			}
		},
	})

	unit.MovementHandler.moveSpell = unit.GetOrRegisterSpell(SpellConfig{
		ActionID: ActionID{OtherID: proto.OtherAction_OtherActionMove},
		Flags:    SpellFlagMeleeMetrics,
		ExtraCastCondition: func(_ *Simulation, _ *Unit) bool {
			return !unit.IsRooted() && !unit.PseudoStats.Stunned
		},

		ApplyEffects: func(sim *Simulation, target *Unit, spell *Spell) {
			unit.MovementHandler.moveAura.Activate(sim)
			unit.MovementHandler.moveAura.SetStacks(sim, max(1, int32(unit.DistanceFromTarget)))
		},
	})
}

func (unit *Unit) IsMoving() bool {
	return unit.MovementHandler.Moving
}

func (unit *Unit) MoveTo(moveRange float64, sim *Simulation) {
	if unit.IsRooted() || unit.PseudoStats.Stunned || moveRange == unit.DistanceFromTarget {
		return
	}
	move := unit.MovementHandler
	move.updateProgress(sim)
	if !move.moveSpell.Cast(sim, unit.CurrentTarget) {
		return
	}
	move.moveRange, move.lastUpdated = moveRange, sim.CurrentTime
	unit.UpdateMovementControl(sim)
}

func (move *MovementHandler) updateProgress(sim *Simulation) {
	if !move.moveAura.IsActive() || !move.unit.IsEnabled() {
		return
	}
	elapsed := max(time.Duration(0), sim.CurrentTime-move.lastUpdated).Seconds()
	move.lastUpdated = sim.CurrentTime
	remaining := move.moveRange - move.unit.DistanceFromTarget
	travelled := min(math.Abs(remaining), elapsed*move.activeSpeed)
	move.unit.DistanceFromTarget += math.Copysign(travelled, remaining)
	move.moveAura.SetStacks(sim, max(1, int32(move.unit.DistanceFromTarget)))
	if math.Abs(move.moveRange-move.unit.DistanceFromTarget) < 1e-9 {
		move.unit.DistanceFromTarget = move.moveRange
		move.moveAura.Deactivate(sim)
	}
}

// Settle the elapsed distance at the old speed before pausing or rescheduling.
// Stun transitions call this as well; a root expiring must not resume a stun.
func (unit *Unit) UpdateMovementControl(sim *Simulation) {
	move := unit.MovementHandler
	if move == nil {
		return
	}
	move.updateProgress(sim)
	if move.moveAction != nil {
		move.moveAction.Cancel(sim)
		move.moveAction = nil
	}
	if !move.moveAura.IsActive() {
		return
	}
	if !unit.IsEnabled() {
		move.Moving, move.activeSpeed = false, 0
		return
	}
	move.activeSpeed = move.MoveSpeed
	if unit.IsRooted() || unit.PseudoStats.Stunned {
		move.activeSpeed = 0
	}
	wasMoving := move.Moving
	move.Moving = move.activeSpeed > 0
	if !move.Moving {
		// There is deliberately no timer while immobile, nor a division by zero.
		if !unit.PseudoStats.Stunned {
			unit.AutoAttacks.EnableAutoSwing(sim)
		}
		return
	}
	if !wasMoving {
		unit.InterruptCast(sim)
		unit.AutoAttacks.CancelAutoSwing(sim)
	}
	remaining := math.Abs(move.moveRange - unit.DistanceFromTarget)
	interval := max(time.Nanosecond, DurationFromSeconds(min(1, remaining)/move.activeSpeed))
	move.moveAction = &PendingAction{
		NextActionAt: sim.CurrentTime + interval,
		OnAction: func(sim *Simulation) {
			move.moveAction = nil
			unit.UpdateMovementControl(sim)
		},
	}
	sim.AddPendingAction(move.moveAction)
}

func (unit *Unit) IsRooted() bool {
	return unit.MovementHandler != nil && unit.MovementHandler.rootCount > 0
}

func (unit *Unit) AddRoot(sim *Simulation) {
	move := unit.MovementHandler
	move.updateProgress(sim)
	move.rootCount++
	unit.UpdateMovementControl(sim)
}

func (unit *Unit) RemoveRoot(sim *Simulation) {
	move := unit.MovementHandler
	if move.rootCount > 0 {
		move.rootCount--
	}
	unit.UpdateMovementControl(sim)
}

// Dynamic variants are for auras which can change speed during an existing move.
func (unit *Unit) AddMoveSpeedModifierDynamic(sim *Simulation, actionID *ActionID, modifier float64) {
	unit.MovementHandler.updateProgress(sim)
	unit.AddMoveSpeedModifier(actionID, modifier)
	unit.UpdateMovementControl(sim)
}

func (unit *Unit) RemoveMoveSpeedModifierDynamic(sim *Simulation, actionID *ActionID) {
	unit.MovementHandler.updateProgress(sim)
	unit.RemoveMoveSpeedModifier(actionID)
	unit.UpdateMovementControl(sim)
}

// A move speed increase of 30% should be represented as 1.30 and a move speed slow of 70% should be respresented as 0.70
func (unit *Unit) AddMoveSpeedModifier(actionId *ActionID, modifier float64) {
	moveSpeedMod := MoveModifier{
		ActionId: actionId,
		Modifier: modifier,
	}
	if moveSpeedMod.Modifier < 1 {
		unit.MovementHandler.addMoveSpeedModifier(unit.MovementHandler.moveSpeedPenalties, moveSpeedMod)
	} else {
		unit.MovementHandler.addMoveSpeedModifier(unit.MovementHandler.moveSpeedBonuses, moveSpeedMod)
	}

}

func (unit *Unit) RemoveMoveSpeedModifier(actionID *ActionID) {
	unit.MovementHandler.removeMoveSpeedModifier(unit.MovementHandler.moveSpeedPenalties, actionID)
	unit.MovementHandler.removeMoveSpeedModifier(unit.MovementHandler.moveSpeedBonuses, actionID)
}
