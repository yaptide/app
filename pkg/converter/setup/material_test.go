package setup

import (
	"testing"

	test "github.com/yaptide/yaptide/pkg/converter/test"
)

var materialTestCasses = test.MarshallingCases{
	{
		&Material{MaterialID(1),
			MaterialPredefined{PredefinedID: "methanol"},
		},
		`{
			"id": 1,
			"specs": {
				"type": "predefined",
				"predefinedId": "methanol"
			}
		}`,
	},
	{
		&Material{MaterialID(1), MaterialPredefined{
			PredefinedID:              "methanol",
			StateOfMatter:             Liquid,
			Density:                   0.001,
			LoadExternalStoppingPower: false,
		}},
		`{
			"id": 1,
			"specs": {
				"type": "predefined",
				"predefinedId": "methanol",
				"density": 0.001,
				"stateOfMatter": "liquid"
			}
		}`,
	},
	{
		&Material{MaterialID(1), MaterialCompound{
			Name:          "ala",
			Density:       1.2345,
			StateOfMatter: Gas,
			Elements: []Element{
				Element{Isotope: "As-75", RelativeStoichiometricFraction: 1},
				Element{Isotope: "H-1 - Hydrogen", RelativeStoichiometricFraction: 8},
			},
		}},
		`{
			"id": 1,
			"specs": {
				"type": "compound",
				"name": "ala",
				"density": 1.2345,
				"stateOfMatter": "gas",
				"elements": [
					{
						"isotope": "As-75",
						"relativeStoichiometricFraction": 1
					},
					{
						"isotope": "H-1 - Hydrogen",
						"relativeStoichiometricFraction": 8
					}
				]
			}
		}`,
	},
	{
		&Material{MaterialID(1), MaterialCompound{
			Name:          "kot",
			Density:       99.9,
			StateOfMatter: Liquid,
			Elements: []Element{
				Element{Isotope: "Gd-*", RelativeStoichiometricFraction: 2, AtomicMass: func() *int64 { i := int64(100); return &i }()},
				Element{Isotope: "U-235", RelativeStoichiometricFraction: 123, IValue: func() *float64 { i := float64(555.34); return &i }()},
			},
			ExternalStoppingPowerFromPredefined: "Water",
		}},
		`{
			"id": 1,
			"specs": {
				"type": "compound",
				"name": "kot",
				"density": 99.9,
				"stateOfMatter": "liquid",
				"elements": [
					{
						"isotope": "Gd-*",
						"relativeStoichiometricFraction": 2,
						"atomicMass": 100
					},
					{
						"isotope": "U-235",
						"relativeStoichiometricFraction": 123,
						"iValue": 555.34
					}
				],
				"externalStoppingPowerFromPredefined": "Water"
			}
		}`,
	},
}

func TestMaterialMarshal(t *testing.T) {
	test.Marshal(t, materialTestCasses)
}

func TestMaterialUnmarshal(t *testing.T) {
	test.Unmarshal(t, materialTestCasses)
}

func TestMaterialUnmarshalMarshalled(t *testing.T) {
	test.UnmarshalMarshalled(t, materialTestCasses)
}

func TestMaterialMarshalUnmarshalled(t *testing.T) {
	test.MarshalUnmarshalled(t, materialTestCasses)
}
