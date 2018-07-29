package simulation

import (
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
	},
}

func TestSerializeBeam(t *testing.T) {
	test.RunSerializeTestCases(t, beamTestCases, beamMarshaler, beamUnmarshaler)
}
