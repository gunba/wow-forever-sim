package core

// AddStun suspends actions without disabling the unit or removing its auras.
// Callers own immunity/eligibility and pair each acquisition with RemoveStun.
func (unit *Unit) AddStun(sim *Simulation) {
	unit.stunCount++
	if unit.stunCount != 1 {
		return
	}
	unit.stunPriorFlag = unit.PseudoStats.Stunned
	unit.stunResumeAutos = unit.AutoAttacks.enabled
	unit.PseudoStats.Stunned = true
	unit.UpdateMovementControl(sim)
	unit.InterruptCast(sim)
	unit.AutoAttacks.CancelAutoSwing(sim)
	if unit.gcdAction != nil {
		unit.CancelGCDTimer(sim)
	}
}

func (unit *Unit) RemoveStun(sim *Simulation) {
	if unit.stunCount == 0 {
		return
	}
	unit.stunCount--
	if unit.stunCount != 0 {
		return
	}
	unit.PseudoStats.Stunned = unit.stunPriorFlag
	unit.UpdateMovementControl(sim)
	resumeAutos := unit.stunResumeAutos
	unit.stunResumeAutos = false
	if !unit.IsEnabled() || unit.PseudoStats.Stunned || sim.CurrentTime >= sim.Duration {
		return
	}
	if resumeAutos && (unit.MovementHandler == nil || !unit.IsMoving()) {
		unit.AutoAttacks.EnableAutoSwing(sim)
	}
	if unit.GCD != nil {
		unit.SetGCDTimer(sim, max(sim.CurrentTime, unit.GCD.ReadyAt()))
	}
}

// InterruptCast cancels a hardcast/channel without spending its completion cost.
// It preserves any unexpired original GCD; cancelling a long cast does not leave
// the unit locked until the former cast completion time.
func (unit *Unit) InterruptCast(sim *Simulation) {
	casting := unit.IsCasting(sim)
	channeling := unit.IsChanneling(sim)
	if !casting && !channeling {
		return
	}
	hc := unit.Hardcast
	gcdReadyAt := hc.GCDReadyAt
	if channeling && !casting && unit.GCD != nil {
		gcdReadyAt = unit.GCD.ReadyAt()
	}
	unit.Hardcast = Hardcast{Expires: startingCDTime}
	if unit.hardcastAction != nil {
		unit.hardcastAction.Cancel(sim)
		unit.hardcastAction = nil
	}
	if channeling {
		unit.ChanneledDot.Cancel(sim)
	}
	if hc.OnInterrupt != nil {
		hc.OnInterrupt(sim)
	}
	if unit.GCD != nil {
		unit.SetGCDTimer(sim, max(sim.CurrentTime, gcdReadyAt))
	}
}
