package material

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

const matTc1Expected = `MEDIUM 1
ICRU 32
END
MEDIUM 2
STATE 1
RHO 100.000000
ICRU 123
LOADDEDX
END
`

const matTc2Expected = `MEDIUM 1
STATE 0
RHO 2.000000
NUCLID 1 20
AMASS 20.000000
IVALUE 30.000000
END
MEDIUM 2
STATE 2
RHO 43.120000
NUCLID 1 1
AMASS 1.000000
IVALUE 43.000000
NUCLID 2 43
AMASS 23.000000
IVALUE 43.000000
END
`

const matTc3Expected = `MEDIUM 1
ICRU 32
END
MEDIUM 2
STATE 1
RHO 100.000000
ICRU 123
LOADDEDX
END
MEDIUM 3
STATE 0
RHO 2.000000
NUCLID 1 20
AMASS 20.000000
IVALUE 30.000000
END
MEDIUM 4
STATE 2
RHO 43.120000
NUCLID 1 1
AMASS 1.000000
IVALUE 43.000000
NUCLID 2 43
AMASS 23.000000
IVALUE 43.000000
END
`

func TestSerializeMat(t *testing.T) {
	type testCase struct {
		Input    Materials
		Expected string
	}

	testCases := []testCase{
		testCase{
			Input: Materials{
				Predefined: []PredefinedMaterial{
					PredefinedMaterial{
						ID:            1,
						ICRUNumber:    32,
						StateOfMatter: -1,
					},
					PredefinedMaterial{
						ID:                        2,
						ICRUNumber:                123,
						StateOfMatter:             1,
						Density:                   100.0,
						LoadExternalStoppingPower: true,
					},
				},
				Compound: []CompoundMaterial{},
			},
			Expected: matTc1Expected,
		},
		testCase{
			Input: Materials{
				Predefined: []PredefinedMaterial{},
				Compound: []CompoundMaterial{
					CompoundMaterial{
						ID:            1,
						StateOfMatter: 0,
						Density:       2.0,
						ExternalStoppingPowerFromMaterialICRU: 0,
						Elements: []Element{
							Element{
								ID: 1,
								RelativeStoichiometricFraction: 20,
								AtomicMass:                     20.0,
								IValue:                         30.0,
							},
						},
					},
					CompoundMaterial{
						ID:            2,
						StateOfMatter: 2,
						Density:       43.12,
						ExternalStoppingPowerFromMaterialICRU: 1,
						Elements: []Element{
							Element{
								ID: 1,
								RelativeStoichiometricFraction: 1,
								AtomicMass:                     1.0,
								IValue:                         43.0,
							},
							Element{
								ID: 2,
								RelativeStoichiometricFraction: 43,
								AtomicMass:                     23.0,
								IValue:                         43.0,
							},
						},
					},
				},
			},
			Expected: matTc2Expected,
		},
		testCase{
			Input: Materials{
				Predefined: []PredefinedMaterial{
					PredefinedMaterial{
						ID:            1,
						ICRUNumber:    32,
						StateOfMatter: -1,
					},
					PredefinedMaterial{
						ID:                        2,
						ICRUNumber:                123,
						StateOfMatter:             1,
						Density:                   100.0,
						LoadExternalStoppingPower: true,
					},
				},
				Compound: []CompoundMaterial{
					CompoundMaterial{
						ID:            3,
						StateOfMatter: 0,
						Density:       2.0,
						ExternalStoppingPowerFromMaterialICRU: 0,
						Elements: []Element{
							Element{
								ID: 1,
								RelativeStoichiometricFraction: 20,
								AtomicMass:                     20.0,
								IValue:                         30.0,
							},
						},
					},
					CompoundMaterial{
						ID:            4,
						StateOfMatter: 2,
						Density:       43.12,
						ExternalStoppingPowerFromMaterialICRU: 1,
						Elements: []Element{
							Element{
								ID: 1,
								RelativeStoichiometricFraction: 1,
								AtomicMass:                     1.0,
								IValue:                         43.0,
							},
							Element{
								ID: 2,
								RelativeStoichiometricFraction: 43,
								AtomicMass:                     23.0,
								IValue:                         43.0,
							},
						},
					},
				},
			},
			Expected: matTc3Expected,
		},
	}

	for n, tc := range testCases {
		t.Run(strconv.Itoa(n), func(t *testing.T) {
			actual := Serialize(tc.Input)
			assert.Equal(t, tc.Expected, actual)
		})
	}

}
