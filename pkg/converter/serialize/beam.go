package serialize

import "github.com/yaptide/yaptide/pkg/converter/setup"

func BeamMarshaler(b setup.Beam) marshaler {
	return StructMarshaler(func(m fieldMarshaler) {
		m("direction",
			StructMarshaler(func(m fieldMarshaler) {
				m("phi", Float64Marshaler(b.Direction.Phi))
				m("theta", Float64Marshaler(b.Direction.Theta))
				m("position", PointMarshaler(b.Direction.Position))
			}),
		)
		m("divergence",
			StructMarshaler(func(m fieldMarshaler) {
				m("sigmaX", Float64Marshaler(b.Divergence.SigmaX))
				m("sigmaY", Float64Marshaler(b.Divergence.SigmaY))
				m("distribution", DistributionMarshaler(b.Divergence.Distribution))
			}),
		)
		m("particle", ParticleMarshaler(b.Particle))
		m("initialBaseEnergy", Float64Marshaler(b.InitialBaseEnergy))
		m("initialEnergySigma", Float64Marshaler(b.InitialEnergySigma))
	})
}

func BeamUnmarshaler(b *setup.Beam) unmarshaler {
	return StructUnmarshaler(func(u fieldUnmarshaler) {
		u("direction",
			StructUnmarshaler(func(u fieldUnmarshaler) {
				u("phi", Float64Unmarshaler(&b.Direction.Phi))
				u("theta", Float64Unmarshaler(&b.Direction.Theta))
				u("position", PointUnmarshaler(&b.Direction.Position))
			}),
		)
		u("divergence",
			StructUnmarshaler(func(u fieldUnmarshaler) {
				u("sigmaX", Float64Unmarshaler(&b.Divergence.SigmaX))
				u("sigmaY", Float64Unmarshaler(&b.Divergence.SigmaY))
				u("distribution", DistributionUnmarshaler(&b.Divergence.Distribution))
			}),
		)
		u("particle", ParticleUnmarshaler(&b.Particle))
		u("initialBaseEnergy", Float64Unmarshaler(&b.InitialBaseEnergy))
		u("initialEnergySigma", Float64Unmarshaler(&b.InitialEnergySigma))
	})
}
