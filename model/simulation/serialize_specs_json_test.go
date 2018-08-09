package simulation

import (
	"fmt"
	"testing"

	"github.com/yaptide/yaptide/pkg/converter/geometry"
	"github.com/yaptide/yaptide/pkg/converter/specs"
	"github.com/yaptide/yaptide/test"
	"gopkg.in/mgo.v2/bson"
)

var simSpecsJsonTestCases = []struct {
	RawValue    []byte
	ObjectValue interface{}
}{
	{
		RawValue: []byte(`
		{
			"id": "111111111111111111111111",
			"userId": "222222222222222222222222",
			"specs": {
				"beam": {
					"direction": {
						"phi": 0,
						"position": {
							"x": 0,
							"y": 0,
							"z": 0
						},
						"theta": 0
					},
					"divergence": {
						"distribution": "",
						"sigmaX": 0,
						"sigmaY": 0
					},
					"initialBaseEnergy": 0,
					"initialEnergySigma": 0,
					"particle": null
				},
				"bodies": [
				],
				"detectors": [
				],
				"materials": [
				],
				"options": {
					"antyparticleCorrectionOn": false,
					"energyStragglingType": "",
					"fastNeutronTransportOn": false,
					"lowEnergyNeutronCutOff": 0,
					"meanEnergyLoss": 0,
					"minEnergyLoss": 0,
					"nuclearReactionsOn": false,
					"numberOfGeneratedParticles": 0,
					"scatteringType": ""
				},
				"zones": [
				]
			}
		}
		`),
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
		RawValue: []byte(`
		{
			"id": "111111111111111111111111",
			"userId": "222222222222222222222222",
			"specs": {
				"beam": {
					"direction": {
						"phi": 0.1,
						"position": {
							"x": 0.3,
							"y": 0.4,
							"z": 0.5
						},
						"theta": 0.2
					},
					"divergence": {
						"distribution": "gaussian",
						"sigmaX": 0.6,
						"sigmaY": 0.7
					},
					"initialBaseEnergy": 0.8,
					"initialEnergySigma": 0.9,
					"particle": {
						"type": "neutron"
					}
				},
				"bodies": [
				{
					"geometry": {
						"center": {
							"x": 1,
							"y": 1,
							"z": 1
						},
						"radius": 10,
						"type": "sphere"
					},
					"id": 0
				},
				{
					"geometry": {
						"center": {
							"x": 1,
							"y": 1,
							"z": 1
						},
						"size": {
							"x": 2,
							"y": 2,
							"z": 2
						},
						"type": "cuboid"
					},
					"id": 1
				},
				{
					"geometry": {
						"center": {
							"x": 1,
							"y": 1,
							"z": 1
						},
						"height": 10,
						"radius": 3,
						"type": "cylinder"
					},
					"id": 2
				}
				],
				"detectors": [
				{
					"geometry": {
						"angle": {
							"max": 3.14,
							"min": 1
						},
						"radius": {
							"max": 100,
							"min": 0
						},
						"slices": {
							"angle": 5,
							"radius": 1,
							"z": 100
						},
						"type": "cylinder",
						"zValue": {
							"max": 200,
							"min": 2
						}
					},
					"id": 0,
					"name": "detector 1",
					"particle": {
						"type": "neutron"
					},
					"scoring": {
						"type": "energy"
					}
				},
				{
					"geometry": {
						"normal": {
							"x": 0,
							"y": 5,
							"z": 5
						},
						"point": {
							"x": 1,
							"y": 1,
							"z": 1
						},
						"type": "plane"
					},
					"id": 1,
					"name": "detector 2",
					"particle": {
						"charge": 50,
						"nucleonsCount": 100,
						"type": "heavy_ion"
					},
					"scoring": {
						"material": 1,
						"type": "dlet"
					}
				}
				],
				"materials": [
				{
					"id": 0,
					"name": "material 1",
					"specs": {
						"density": null,
						"predefinedId": "water",
						"stateOfMatter": "liquid",
						"type": "predefined"
					}
				},
				{
					"id": 1,
					"name": "material 2",
					"specs": {
						"density": 2.3,
						"predefinedId": "water",
						"stateOfMatter": "solid",
						"type": "predefined"
					}
				},
				{
					"id": 2,
					"name": "material 3",
					"specs": {
						"density": 1.1,
						"elements": [
						],
						"stateOfMatter": "solid",
						"type": "compound"
					}
				},
				{
					"id": 3,
					"name": "material 4",
					"specs": {
						"density": 2.1,
						"elements": [
						{
							"atomicMass": null,
							"iValue": null,
							"isotope": "isotope 1",
							"relativeStoichiometricFraction": 5
						},
						{
							"atomicMass": 12,
							"iValue": 1000.1,
							"isotope": "isotope 2",
							"relativeStoichiometricFraction": 5
						}
						],
						"stateOfMatter": "gas",
						"type": "compound"
					}
				}
				],
				"options": {
					"antyparticleCorrectionOn": true,
					"energyStragglingType": "vavilov",
					"fastNeutronTransportOn": true,
					"lowEnergyNeutronCutOff": 0.3,
					"meanEnergyLoss": 0.1,
					"minEnergyLoss": 0.2,
					"nuclearReactionsOn": true,
					"numberOfGeneratedParticles": 4,
					"scatteringType": "gaussian"
				},
				"zones": [
				{
					"baseId": 1,
					"color": {
						"a": 100,
						"b": 200,
						"g": 5,
						"r": 1
					},
					"construction": [
					{
						"bodyId": 2,
						"type": "subtract"
					},
					{
						"bodyId": 0,
						"type": "union"
					}
					],
					"id": 0,
					"materialId": 1,
					"name": "zone 1",
					"parentId": 0
				},
				{
					"baseId": 1,
					"color": {
						"a": 100,
						"b": 200,
						"g": 5,
						"r": 1
					},
					"construction": [
					],
					"id": 1,
					"materialId": 1,
					"name": "zone 2",
					"parentId": 0
				}
				]
			}
		}
		`),
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

func TestSerializeSimulationJsonSpec(t *testing.T) {
	for index, simCase := range simSpecsJsonTestCases {
		t.Run(
			fmt.Sprintf("TestSerializeSimulationJsonSpec Marshal case %d", index),
			func(t *testing.T) {
				test.RunJsonMarshallTestCase(
					t, simCase.RawValue, simCase.ObjectValue,
				)
			},
		)
		t.Run(
			fmt.Sprintf("TestSerializeSimulationJsonSpec Unmarshal case %d", index),
			func(t *testing.T) {
				test.RunJsonUnmarshallTestCase(
					t, simCase.ObjectValue, simCase.RawValue,
				)
			},
		)
	}
}
