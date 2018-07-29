package simulation

import (
	"github.com/yaptide/yaptide/pkg/converter/specs"
	. "github.com/yaptide/yaptide/pkg/serialize"
)

func beamMarshaler(b specs.Beam) marshaler {
	return StructMarshaler(func(m fieldMarshaler) {
		m("direction",
			StructMarshaler(func(m fieldMarshaler) {
				m("phi", Float64Marshaler(b.Direction.Phi))
				m("theta", Float64Marshaler(b.Direction.Theta))
				m("position", pointMarshaler(b.Direction.Position))
			}),
		)
		m("divergence",
			StructMarshaler(func(m fieldMarshaler) {
				m("sigmaX", Float64Marshaler(b.Divergence.SigmaX))
				m("sigmaY", Float64Marshaler(b.Divergence.SigmaY))
				m("distribution", distributionMarshaler(b.Divergence.Distribution))
			}),
		)
		m("particle", particleMarshaler(b.Particle))
		m("initialBaseEnergy", Float64Marshaler(b.InitialBaseEnergy))
		m("initialEnergySigma", Float64Marshaler(b.InitialEnergySigma))
	})
}

func beamUnmarshaler(b *specs.Beam) unmarshaler {
	return StructUnmarshaler(func(u fieldUnmarshaler) {
		u("direction",
			StructUnmarshaler(func(u fieldUnmarshaler) {
				u("phi", Float64Unmarshaler(&b.Direction.Phi))
				u("theta", Float64Unmarshaler(&b.Direction.Theta))
				u("position", pointUnmarshaler(&b.Direction.Position))
			}),
		)
		u("divergence",
			StructUnmarshaler(func(u fieldUnmarshaler) {
				u("sigmaX", Float64Unmarshaler(&b.Divergence.SigmaX))
				u("sigmaY", Float64Unmarshaler(&b.Divergence.SigmaY))
				u("distribution", distributionUnmarshaler(&b.Divergence.Distribution))
			}),
		)
		u("particle", particleUnmarshaler(&b.Particle))
		u("initialBaseEnergy", Float64Unmarshaler(&b.InitialBaseEnergy))
		u("initialEnergySigma", Float64Unmarshaler(&b.InitialEnergySigma))
	})
}
