package smite

import (
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/priest"
)

func RegisterSmitePriest() {
	core.RegisterAgentFactory(
		proto.Player_SmitePriest{},
		proto.Spec_SpecSmitePriest,
		func(character *core.Character, options *proto.Player) core.Agent {
			return NewSmitePriest(character, options)
		},
		func(player *proto.Player, spec interface{}) {
			playerSpec, ok := spec.(*proto.Player_SmitePriest)
			if !ok {
				panic("Invalid spec value for Smite Priest!")
			}
			player.Spec = playerSpec
		},
	)
}

func NewSmitePriest(character *core.Character, options *proto.Player) *SmitePriest {
	smiteOptions := options.GetSmitePriest()
	basePriest := priest.New(character, options.TalentsString)
	basePriest.Latency = float64(basePriest.ChannelClipDelay.Milliseconds())
	basePriest.PowerInfusionTarget = smiteOptions.GetOptions().GetPowerInfusionTarget()

	return &SmitePriest{
		Priest:  basePriest,
		options: smiteOptions.Options,
	}
}

type SmitePriest struct {
	*priest.Priest
	options *proto.SmitePriest_Options
}

func (spriest *SmitePriest) GetPriest() *priest.Priest {
	return spriest.Priest
}

func (spriest *SmitePriest) Initialize() {
	spriest.Priest.Initialize()
}

func (spriest *SmitePriest) Reset(_ *core.Simulation) {
}
