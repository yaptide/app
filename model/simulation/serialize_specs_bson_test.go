package simulation

import (
	"fmt"
	"testing"

	"github.com/yaptide/yaptide/pkg/converter/geometry"
	"github.com/yaptide/yaptide/pkg/converter/specs"
	"github.com/yaptide/yaptide/test"
	"gopkg.in/mgo.v2/bson"
)

var simSpecsBsonTestCases = []test.SerializeTestCase{
	{
		RawValue: mRaw{
			"_id":    bson.ObjectIdHex("111111111111111111111111"),
			"userId": bson.ObjectIdHex("222222222222222222222222"),
			"specs": mRaw{
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
		},
		ObjectValue: SpecsDB{
			ID:     bson.ObjectIdHex("111111111111111111111111"),
			UserID: bson.ObjectIdHex("222222222222222222222222"),
			Specs: Specs{
				Materials: []Material{},
				Bodies:    []Body{},
				Zones:     []Zone{},
				Detectors: []Detector{},
			},
		},
	},
	{
		RawValue: mRaw{
			"_id":    bson.ObjectIdHex("111111111111111111111111"),
			"userId": bson.ObjectIdHex("222222222222222222222222"),
			"specs": mRaw{
				"materials": []interface{}{
					mRaw{
						"id":   int64(0),
						"name": "material 1",
						"specs": mRaw{
							"type":          "predefined",
							"predefinedId":  "water",
							"density":       nil,
							"stateOfMatter": "liquid",
						},
					},
					mRaw{
						"id":   int64(1),
						"name": "material 2",
						"specs": mRaw{
							"type":          "predefined",
							"predefinedId":  "water",
							"density":       2.3,
							"stateOfMatter": "solid",
						},
					},
					mRaw{
						"id":   int64(2),
						"name": "material 3",
						"specs": mRaw{
							"type":          "compound",
							"density":       1.1,
							"stateOfMatter": "solid",
							"elements":      aRaw{},
						},
					},
					mRaw{
						"id":   int64(3),
						"name": "material 4",
						"specs": mRaw{
							"type":          "compound",
							"density":       2.1,
							"stateOfMatter": "gas",
							"elements": aRaw{
								mRaw{
									"isotope":                        "isotope 1",
									"relativeStoichiometricFraction": int64(5),
									"atomicMass":                     nil,
									"iValue":                         nil,
								},
								mRaw{
									"isotope":                        "isotope 2",
									"relativeStoichiometricFraction": int64(5),
									"atomicMass":                     int64(12),
									"iValue":                         1000.1,
								},
							},
						},
					},
				},
				"bodies": []interface{}{
					mRaw{
						"id": int64(0),
						"geometry": mRaw{
							"type":   "sphere",
							"center": mRaw{"x": float64(1), "y": float64(1), "z": float64(1)},
							"radius": float64(10),
						},
					},
					mRaw{
						"id": int64(1),
						"geometry": mRaw{
							"type":   "cuboid",
							"center": mRaw{"x": float64(1), "y": float64(1), "z": float64(1)},
							"size":   mRaw{"x": float64(2), "y": float64(2), "z": float64(2)},
						},
					},
					mRaw{
						"id": int64(2),
						"geometry": mRaw{
							"type":   "cylinder",
							"center": mRaw{"x": float64(1), "y": float64(1), "z": float64(1)},
							"height": float64(10),
							"radius": float64(3),
						},
					},
				},
				"zones": []interface{}{
					mRaw{
						"name": "zone 1",
						"color": mRaw{
							"r": int64(1), "g": int64(5), "b": int64(200), "a": int64(100),
						},
						"id":         int64(0),
						"parentId":   int64(0),
						"baseId":     int64(1),
						"materialId": int64(1),
						"construction": aRaw{
							mRaw{"type": "subtract", "bodyId": int64(2)},
							mRaw{"type": "union", "bodyId": int64(0)},
						},
					},
					mRaw{
						"name": "zone 2",
						"color": mRaw{
							"r": int64(1), "g": int64(5), "b": int64(200), "a": int64(100),
						},
						"id":           int64(1),
						"parentId":     int64(0),
						"baseId":       int64(1),
						"materialId":   int64(1),
						"construction": aRaw{},
					},
				},
				"detectors": []interface{}{
					mRaw{
						"name": "detector 1",
						"id":   int64(0),
						"geometry": mRaw{
							"type":   "cylinder",
							"radius": mRaw{"min": float64(0), "max": float64(100)},
							"angle":  mRaw{"min": float64(1), "max": float64(3.14)},
							"zValue": mRaw{"min": float64(2), "max": float64(200)},
							"slices": mRaw{
								"radius": int64(1), "angle": int64(5), "z": int64(100),
							},
						},
						"particle": mRaw{"type": "neutron"},
						"scoring":  mRaw{"type": "energy"},
					},
					mRaw{
						"name": "detector 2",
						"id":   int64(1),
						"geometry": mRaw{
							"type": "plane",
							"point": mRaw{
								"x": float64(1), "y": float64(1), "z": float64(1),
							},
							"normal": mRaw{
								"x": float64(0), "y": float64(5), "z": float64(5),
							},
						},
						"particle": mRaw{
							"type":          "heavy_ion",
							"charge":        int64(50),
							"nucleonsCount": int64(100),
						},
						"scoring": mRaw{"type": "dlet", "material": int64(1)},
					},
				},
				"beam": mRaw{
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
				"options": mRaw{
					"antyparticleCorrectionOn":   true,
					"nuclearReactionsOn":         true,
					"meanEnergyLoss":             float64(0.1),
					"minEnergyLoss":              float64(0.2),
					"scatteringType":             "gaussian",
					"energyStragglingType":       "vavilov",
					"fastNeutronTransportOn":     true,
					"lowEnergyNeutronCutOff":     float64(0.3),
					"numberOfGeneratedParticles": int64(4),
				},
			},
		},
		ObjectValue: SpecsDB{
			ID:     bson.ObjectIdHex("111111111111111111111111"),
			UserID: bson.ObjectIdHex("222222222222222222222222"),
			Specs: Specs{
				Materials: []Material{
					{
						Name: "material 1",
						Material: specs.Material{
							ID:    specs.MaterialID(0),
							Specs: specs.MaterialPredefined{"water", nil, specs.Liquid},
						},
					},
					{
						Name: "material 2",
						Material: specs.Material{
							ID: specs.MaterialID(1),
							Specs: specs.MaterialPredefined{
								"water",
								func() *float64 { a := 2.3; return &a }(),
								specs.Solid,
							},
						},
					}, {
						Name: "material 3",
						Material: specs.Material{
							ID:    specs.MaterialID(2),
							Specs: specs.MaterialCompound{1.1, specs.Solid, []specs.Element{}},
						},
					}, {
						Name: "material 4",
						Material: specs.Material{
							ID: specs.MaterialID(3),
							Specs: specs.MaterialCompound{
								2.1,
								specs.Gas,
								[]specs.Element{
									{"isotope 1", 5, nil, nil},
									{
										"isotope 2",
										5,
										func() *int64 { a := int64(12); return &a }(),
										func() *float64 { a := 1000.1; return &a }(),
									},
								},
							},
						},
					},
				},
				Bodies: []Body{
					{
						ID: specs.BodyID(0),
						Geometry: specs.BodySphere{
							Center: geometry.Point{X: 1, Y: 1, Z: 1},
							Radius: 10,
						},
					}, {
						ID: specs.BodyID(1),
						Geometry: specs.BodyCuboid{
							Center: geometry.Point{X: 1, Y: 1, Z: 1},
							Size:   geometry.Vec3D{X: 2, Y: 2, Z: 2},
						},
					}, {
						ID: specs.BodyID(2),
						Geometry: specs.BodyCylinder{
							Center: geometry.Point{X: 1, Y: 1, Z: 1},
							Height: 10,
							Radius: 3,
						},
					},
				},
				Zones: []Zone{
					{
						Name:  "zone 1",
						Color: Color{R: 1, G: 5, B: 200, A: 100},
						Zone: specs.Zone{
							ID:         specs.ZoneID(0),
							ParentID:   specs.ZoneID(0),
							BaseID:     specs.BodyID(1),
							MaterialID: specs.MaterialID(1),
							Construction: []specs.ZoneOperation{
								{specs.Subtract, specs.BodyID(2)},
								{specs.Union, specs.BodyID(0)},
							},
						},
					}, {
						Name:  "zone 2",
						Color: Color{R: 1, G: 5, B: 200, A: 100},
						Zone: specs.Zone{
							ID:           specs.ZoneID(1),
							ParentID:     specs.ZoneID(0),
							BaseID:       specs.BodyID(1),
							MaterialID:   specs.MaterialID(1),
							Construction: []specs.ZoneOperation{},
						},
					},
				},
				Detectors: []Detector{
					{
						Name: "detector 1",
						Detector: specs.Detector{
							ID: specs.DetectorID(0),
							Geometry: specs.DetectorCylinder{
								Radius: geometry.Range{0, 100},
								Angle:  geometry.Range{1, 3.14},
								ZValue: geometry.Range{2, 200},
								Slices: geometry.Vec3DCylindricalInt{1, 5, 100},
							},
							ScoredParticle: specs.PredefinedParticle("neutron"),
							Scoring:        specs.PredefinedScoring("energy"),
						},
					}, {
						Name: "detector 2",
						Detector: specs.Detector{
							ID: specs.DetectorID(1),
							Geometry: specs.DetectorPlane{
								Point:  geometry.Point{1, 1, 1},
								Normal: geometry.Vec3D{0, 5, 5},
							},
							ScoredParticle: specs.HeavyIon{
								Charge:        50,
								NucleonsCount: 100,
							},
							Scoring: specs.LetTypeScoring{
								Type:     "dlet",
								Material: specs.MaterialID(1),
							},
						},
					},
				},
				Beam: Beam{
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
				Options: SimulationOptions{
					AntyparticleCorrectionOn:   true,
					NuclearReactionsOn:         true,
					MeanEnergyLoss:             0.1,
					MinEnergyLoss:              0.2,
					ScatteringType:             specs.GaussianScattering,
					EnergyStraggling:           specs.VavilovStraggling,
					FastNeutronTransportOn:     true,
					LowEnergyNeutronCutOff:     0.3,
					NumberOfGeneratedParticles: 4,
				},
			},
		},
	},
}

func TestSerializeSimulationBsonSpec(t *testing.T) {
	for index, simCase := range simSpecsBsonTestCases {
		t.Run(
			fmt.Sprintf("TestSerializeSimulationBsonSpec Marshal case %d", index),
			func(t *testing.T) {
				test.RunBsonMarshallTestCase(
					t, simCase.RawValue, simCase.ObjectValue,
				)
			},
		)
		t.Run(
			fmt.Sprintf("TestSerializeSimulationBsonSpec Unmarshal case %d", index),
			func(t *testing.T) {
				test.RunBsonUnmarshallTestCase(
					t, simCase.ObjectValue, simCase.RawValue,
				)
			},
		)
	}
}
