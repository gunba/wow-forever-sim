package hunter

import (
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

func (hunter *Hunter) NewAPLValue(_ *core.APLRotation, config *proto.APLValue) core.APLValue {
	switch config.Value.(type) {
	case *proto.APLValue_HunterHawkCount:
		return &APLValueHunterHawkCount{hunter: hunter}
	case *proto.APLValue_HunterHawkRemainingTime:
		return &APLValueHunterHawkRemainingTime{hunter: hunter}
	default:
		return nil
	}
}

// Guardian state is owner-global, not a victim-bound aura or DoT. In-flight
// summons are not active, and the earliest expiry is independent of targeting.
func (hunter *Hunter) activeHawkState(sim *core.Simulation) (int32, time.Duration) {
	count := int32(0)
	earliest := core.NeverExpires
	for _, hawk := range hunter.Hawks {
		if hawk == nil || !hawk.IsEnabled() || hawk.expiresAt <= sim.CurrentTime {
			continue
		}
		count++
		earliest = min(earliest, hawk.expiresAt)
	}
	if count == 0 {
		return 0, 0
	}
	return count, earliest - sim.CurrentTime
}

type APLValueHunterHawkCount struct {
	core.DefaultAPLValueImpl
	hunter *Hunter
}

func (*APLValueHunterHawkCount) Type() proto.APLValueType {
	return proto.APLValueType_ValueTypeInt
}
func (value *APLValueHunterHawkCount) GetInt(sim *core.Simulation) int32 {
	count, _ := value.hunter.activeHawkState(sim)
	return count
}
func (*APLValueHunterHawkCount) String() string { return "Hunter Hawk Count" }

type APLValueHunterHawkRemainingTime struct {
	core.DefaultAPLValueImpl
	hunter *Hunter
}

func (*APLValueHunterHawkRemainingTime) Type() proto.APLValueType {
	return proto.APLValueType_ValueTypeDuration
}
func (value *APLValueHunterHawkRemainingTime) GetDuration(sim *core.Simulation) time.Duration {
	_, remaining := value.hunter.activeHawkState(sim)
	return remaining
}
func (*APLValueHunterHawkRemainingTime) String() string { return "Hunter Hawk Remaining Time" }
