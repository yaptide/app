package simulation

import (
	"github.com/yaptide/yaptide/pkg/converter/geometry"
	"github.com/yaptide/yaptide/pkg/converter/specs"
	"github.com/yaptide/yaptide/test"
	"testing"
)

var beamTestCases = []test.SerializeTestCase{
	{
		RawValue: mRaw{
			"direction": mRaw{
				"phi":   float64(0),
				"theta": float64(0),
				"position": mRaw{
					"x": float64(0),
					"y": float64(0),
					"z": float64(0),
				},
			},
			"divergence": mRaw{
				"sigmaX":       float64(0),
				"sigmaY":       float64(0),
				"distribution": "",
			},
			"particle":           nil,
			"initialBaseEnergy":  float64(0),
			"initialEnergySigma": float64(0),
		},
		ObjectValue: Beam{},
	}, {
		RawValue: mRaw{
			"direction": mRaw{
				"phi":   float64(0.1),
				"theta": float64(0.2),
				"position": mRaw{
					"x": float64(0.3),
					"y": float64(0.4),
					"z": float64(0.5),
				},
			},
			"divergence": mRaw{
				"sigmaX":       float64(0.6),
				"sigmaY":       float64(0.7),
				"distribution": "gaussian",
			},
			"particle": mRaw{
				"type": "neutron",
			},
			"initialBaseEnergy":  float64(0.8),
			"initialEnergySigma": float64(0.9),
		},
		ObjectValue: Beam{
			Direction: specs.BeamDirection{
				0.1, 0.2, geometry.Point{0.3, 0.4, 0.5},
			},
			Divergence: specs.BeamDivergence{
				0.6, 0.7, specs.GaussianDistribution,
			},
			Particle:           specs.PredefinedParticle("neutron"),
			InitialBaseEnergy:  0.8,
			InitialEnergySigma: 0.9,
		},
	},
}

func TestSerializeBeam(t *testing.T) {
	test.RunSerializeTestCases(t, beamTestCases, beamMarshaler, beamUnmarshaler)
}
