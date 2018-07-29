package simulation

import (
	"testing"

	"github.com/yaptide/yaptide/test"
)

var simSpecsTestCases = []test.SerializeTestCase{
	{
		RawValue: mRaw{
			"materials": []interface{}{},
			"bodies":    []interface{}{},
			"zones":     []interface{}{},
			"detectors": []interface{}{},
			"beam": mRaw{
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
			"options": mRaw{
				"antyparticleCorrectionOn":   false,
				"nuclearReactionsOn":         false,
				"meanEnergyLoss":             float64(0),
				"minEnergyLoss":              float64(0),
				"scatteringType":             "",
				"energyStragglingType":       "",
				"fastNeutronTransportOn":     false,
				"lowEnergyNeutronCutOff":     float64(0),
				"numberOfGeneratedParticles": int64(0),
			},
		},
		ObjectValue: Specs{
			Materials: []Material{},
			Bodies:    []Body{},
			Zones:     []Zone{},
			Detectors: []Detector{},
		},
	},
}

func TestSerializeSimulationSpec(t *testing.T) {
	test.RunSerializeTestCases(
		t, simSpecsTestCases,
		specsMarshaler, specsUnmarshaler,
	)
}
