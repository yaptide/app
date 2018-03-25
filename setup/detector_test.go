package setup

import (
	"testing"

	"github.com/yaptide/converter/common"
	test "github.com/yaptide/converter/test"
)

var detectorTestCasses = test.MarshallingCases{
	{
		&Detector{
			ID:   ID(1),
			Name: "ala",
			DetectorGeometry: Mesh{
				Center: common.Point{X: 1, Y: 2, Z: 3},
				Size:   common.Vec3D{X: 1, Y: 2, Z: 3},
				Slices: common.Vec3DInt{X: 10, Y: 10, Z: 10},
			},
			ScoredParticle: common.PredefinedParticle("all"),
			ScoringType:    PredefinedScoring("energy"),
		},
		`{
			"id": 1,
			"name": "ala",
			"detectorGeometry": {
				"type": "mesh",
				"center": {
					"x": 1,
					"y": 2,
					"z": 3
				},
				"size": {
					"x": 1,
					"y": 2,
					"z": 3
				},
				"slices": {
					"x": 10,
					"y": 10,
					"z": 10
				}
			},
			"particle": {
				"type": "all"
			},
			"scoring": {
				"type": "energy"
			}
		}`,
	}, {
		&Detector{
			ID:   ID(1),
			Name: "ma",
			DetectorGeometry: Mesh{
				Center: common.Point{X: 1, Y: 2, Z: 3},
				Size:   common.Vec3D{X: 1, Y: 2, Z: 3},
				Slices: common.Vec3DInt{X: 10, Y: 10, Z: 10},
			},
			ScoredParticle: common.HeavyIon{Charge: 10, NucleonsCount: 10},
			ScoringType:    LetTypeScoring{Type: "tlet", Material: 0},
		},
		`{
			"id": 1,
			"name": "ma",
			"detectorGeometry": {
				"type": "mesh",
				"center": {
					"x": 1,
					"y": 2,
					"z": 3
				},
				"size": {
					"x": 1,
					"y": 2,
					"z": 3
				},
				"slices": {
					"x": 10,
					"y": 10,
					"z": 10
				}
			},
			"particle": {
				"type": "heavy_ion",
				"charge": 10,
				"nucleonsCount": 10
			},
			"scoring": {
				"type": "tlet",
				"material": 0
			}
		}`,
	},
}

func TestDetectorMarshal(t *testing.T) {
	test.Marshal(t, detectorTestCasses)
}

func TestDetectorUnmarshal(t *testing.T) {
	test.Unmarshal(t, detectorTestCasses)
}

func TestDetectorUnmarshalMarshalled(t *testing.T) {
	test.UnmarshalMarshalled(t, detectorTestCasses)
}

func TestDetectorMarshalUnmarshalled(t *testing.T) {
	test.MarshalUnmarshalled(t, detectorTestCasses)
}
