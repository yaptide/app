package material

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yaptide/yaptide/pkg/converter/shield/mapping"
	"github.com/yaptide/yaptide/pkg/converter/specs"
)

func TestSuccessfullMaterialsConvert(t *testing.T) {
	type testCase struct {
		Input                      map[specs.MaterialID]specs.Material
		Expected                   Materials
		ExpectedMaterialIDToShield map[specs.MaterialID]ShieldID
	}

	check := func(t *testing.T, tc testCase) {
		t.Helper()

		actual, actualMaterialIDToShield, actualErr := ConvertMaterials(tc.Input)

		assert.Equal(t, nil, actualErr)
		assert.Equal(t, tc.Expected, actual)
		assert.Equal(t, tc.ExpectedMaterialIDToShield, actualMaterialIDToShield)
	}

	convertedSimplePredefined := PredefinedMaterial{
		ICRUNumber:    mapping.MaterialICRU(273),
		StateOfMatter: mapping.StateNonDefined,
	}

	convertedFullPredefined := PredefinedMaterial{
		ICRUNumber:    mapping.MaterialICRU(198),
		StateOfMatter: mapping.StateLiquid,
		Density:       func() *float64 { a := 123.45; return &a }(),
	}

	convertedCompound := CompoundMaterial{
		StateOfMatter: mapping.StateSolid,
		Density:       99.9,
		Elements: []Element{
			Element{
				ID: 64,
				RelativeStoichiometricFraction: 2,
				AtomicMass:                     func() *int64 { a := int64(100); return &a }(),
				IValue:                         func() *float64 { a := float64(0.0); return &a }(),
			},
			Element{
				ID: 103,
				RelativeStoichiometricFraction: 123,
				AtomicMass:                     func() *int64 { a := int64(0.0); return &a }(),
				IValue:                         func() *float64 { a := float64(555.34); return &a }(),
			},
		},
	}

	convertedAnotherCompound := CompoundMaterial{
		StateOfMatter: mapping.StateGas,
		Density:       0.999,
		Elements: []Element{
			Element{
				ID: 6,
				RelativeStoichiometricFraction: 4,
				AtomicMass:                     func() *int64 { a := int64(1); return &a }(),
				IValue:                         func() *float64 { a := float64(0.0); return &a }(),
			},
			Element{
				ID: 14,
				RelativeStoichiometricFraction: 1,
				AtomicMass:                     func() *int64 { a := int64(0); return &a }(),
				IValue:                         func() *float64 { a := float64(0.34); return &a }(),
			},
			Element{
				ID: 11,
				RelativeStoichiometricFraction: 11111,
				AtomicMass:                     func() *int64 { a := int64(987); return &a }(),
				IValue:                         func() *float64 { a := float64(0.123); return &a }(),
			},
		},
	}

	t.Run("OnePredefined", func(t *testing.T) {
		check(t,
			testCase{
				Input: createMaterialMap(genSimplePredefined(1)),
				Expected: Materials{
					Predefined: []PredefinedMaterial{setPredefinedID(convertedSimplePredefined, 1)},
					Compound:   []CompoundMaterial{},
				},
				ExpectedMaterialIDToShield: map[specs.MaterialID]ShieldID{1: 1},
			})

		check(t,
			testCase{
				Input: createMaterialMap(genFullPredefined(6)),
				Expected: Materials{
					Predefined: []PredefinedMaterial{setPredefinedID(convertedFullPredefined, 1)},
					Compound:   []CompoundMaterial{},
				},
				ExpectedMaterialIDToShield: map[specs.MaterialID]ShieldID{6: 1},
			},
		)
	})

	t.Run("OneCompound", func(t *testing.T) {
		check(t,
			testCase{
				Input: createMaterialMap(genCompound(1001)),
				Expected: Materials{
					Predefined: []PredefinedMaterial{},
					Compound:   []CompoundMaterial{setCompoundID(convertedCompound, 1)},
				},
				ExpectedMaterialIDToShield: map[specs.MaterialID]ShieldID{1001: 1},
			})

		check(t,
			testCase{
				Input: createMaterialMap(genAnotherCompound(4000)),
				Expected: Materials{
					Predefined: []PredefinedMaterial{},
					Compound:   []CompoundMaterial{setCompoundID(convertedAnotherCompound, 1)},
				},
				ExpectedMaterialIDToShield: map[specs.MaterialID]ShieldID{4000: 1},
			})
	})

	t.Run("FewPredefined", func(t *testing.T) {
		check(t,
			testCase{
				Input: createMaterialMap(
					genSimplePredefined(1),
					genFullPredefined(2)),
				Expected: Materials{
					Predefined: []PredefinedMaterial{
						setPredefinedID(convertedSimplePredefined, 1),
						setPredefinedID(convertedFullPredefined, 2),
					},
					Compound: []CompoundMaterial{},
				},
				ExpectedMaterialIDToShield: map[specs.MaterialID]ShieldID{1: 1, 2: 2},
			})

		check(t,
			testCase{
				Input: createMaterialMap(
					genSimplePredefined(2),
					genFullPredefined(1)),
				Expected: Materials{
					Predefined: []PredefinedMaterial{
						setPredefinedID(convertedFullPredefined, 1),
						setPredefinedID(convertedSimplePredefined, 2),
					},
					Compound: []CompoundMaterial{},
				},
				ExpectedMaterialIDToShield: map[specs.MaterialID]ShieldID{1: 1, 2: 2},
			})
	})

	t.Run("FewCompound", func(t *testing.T) {
		check(t,
			testCase{
				Input: createMaterialMap(
					genCompound(1),
					genAnotherCompound(2)),
				Expected: Materials{
					Predefined: []PredefinedMaterial{},
					Compound: []CompoundMaterial{
						setCompoundID(convertedCompound, 1),
						setCompoundID(convertedAnotherCompound, 2),
					},
				},
				ExpectedMaterialIDToShield: map[specs.MaterialID]ShieldID{1: 1, 2: 2},
			})
		check(t,
			testCase{
				Input: createMaterialMap(
					genCompound(2),
					genAnotherCompound(1)),
				Expected: Materials{
					Predefined: []PredefinedMaterial{},
					Compound: []CompoundMaterial{
						setCompoundID(convertedAnotherCompound, 1),
						setCompoundID(convertedCompound, 2),
					},
				},
				ExpectedMaterialIDToShield: map[specs.MaterialID]ShieldID{1: 1, 2: 2},
			})

	})

	t.Run("Mixed", func(t *testing.T) {
		check(t,
			testCase{
				Input: createMaterialMap(
					genSimplePredefined(1),
					genFullPredefined(2),
					genCompound(3),
					genAnotherCompound(4)),
				Expected: Materials{
					Predefined: []PredefinedMaterial{
						setPredefinedID(convertedSimplePredefined, 1),
						setPredefinedID(convertedFullPredefined, 2),
					},
					Compound: []CompoundMaterial{
						setCompoundID(convertedCompound, 3),
						setCompoundID(convertedAnotherCompound, 4),
					},
				},
				ExpectedMaterialIDToShield: map[specs.MaterialID]ShieldID{1: 1, 2: 2, 3: 3, 4: 4},
			})

		check(t,
			testCase{
				Input: createMaterialMap(
					genSimplePredefined(9),
					genFullPredefined(2),
					genCompound(100),
					genAnotherCompound(3),
					genSimplePredefined(1),
				),
				Expected: Materials{
					Predefined: []PredefinedMaterial{
						setPredefinedID(convertedSimplePredefined, 1),
						setPredefinedID(convertedFullPredefined, 2),
						setPredefinedID(convertedSimplePredefined, 3),
					},
					Compound: []CompoundMaterial{
						setCompoundID(convertedAnotherCompound, 4),
						setCompoundID(convertedCompound, 5),
					},
				},
				ExpectedMaterialIDToShield: map[specs.MaterialID]ShieldID{1: 1, 2: 2, 9: 3, 3: 4, 100: 5},
			})
	})

	t.Run("VacuumShouldBeNotSerialized", func(t *testing.T) {
		check(t,
			testCase{
				Input: createMaterialMap(genVacuum(1)),
				Expected: Materials{
					Predefined: []PredefinedMaterial{},
					Compound:   []CompoundMaterial{},
				},
				ExpectedMaterialIDToShield: map[specs.MaterialID]ShieldID{1: 1000},
			})

		check(t,
			testCase{
				Input: createMaterialMap(
					genSimplePredefined(1),
					genVacuum(2),
					genAnotherCompound(3)),
				Expected: Materials{
					Predefined: []PredefinedMaterial{
						setPredefinedID(convertedSimplePredefined, 1),
					},
					Compound: []CompoundMaterial{
						setCompoundID(convertedAnotherCompound, 2),
					},
				},
				ExpectedMaterialIDToShield: map[specs.MaterialID]ShieldID{1: 1, 2: 1000, 3: 2},
			})

	})

}

func TestBadInputMaterialsConvert(t *testing.T) {
	type testCase struct {
		Input         map[specs.MaterialID]specs.Material
		ExpectedError error
	}

	check := func(t *testing.T, tc testCase) {
		t.Helper()

		actual, _, actualErr := ConvertMaterials(tc.Input)

		assert.Equal(t, Materials{}, actual)
		assert.Equal(t, tc.ExpectedError, actualErr)
	}

	t.Run("ToManyMaterials", func(t *testing.T) {
		const materialsN = 1000
		materials := map[specs.MaterialID]specs.Material{}
		for i := int64(0); i < materialsN; i++ {
			materials[specs.MaterialID(i)] = genSimplePredefined(i)
		}

		check(t, testCase{
			Input: materials,
			ExpectedError: errors.New(
				"[serializer] mat.dat: Only 100 distinct materials" +
					" are permitted in shield (1000 > 100)",
			),
		})
	})

	t.Run("VoxelNotImplemented", func(t *testing.T) {
		const materialsN = 1000
		materials := map[specs.MaterialID]specs.Material{}
		for i := int64(0); i < materialsN; i++ {
			materials[specs.MaterialID(i)] = genSimplePredefined(i)
		}

		check(t, testCase{
			Input: map[specs.MaterialID]specs.Material{
				specs.MaterialID(1): specs.Material{
					ID:    1,
					Specs: specs.MaterialVoxel{},
				},
			},
			ExpectedError: errors.New(
				"[serializer] Material{Id: 1} -> mat.dat: Voxel material" +
					" serialization not implemented",
			),
		})
	})

	t.Run("PredefinedMappingNotFound", func(t *testing.T) {
		const id = 1
		mat := genSimplePredefined(id)
		predef := mat.Specs.(specs.MaterialPredefined)
		predef.PredefinedID = "predefNameNotDefined"
		mat.Specs = predef

		check(t, testCase{
			Input: createMaterialMap(mat),
			ExpectedError: errors.New(
				"[serializer] Material{Id: 1} -> mat.dat: \"predefNameNotDefined\"" +
					" material mapping to shield format not found",
			),
		})
	})

	t.Run("IsotopeMappingNotFound", func(t *testing.T) {
		const id = 1
		mat := genCompound(id)
		compound := mat.Specs.(specs.MaterialCompound)
		compound.Elements[0].Isotope = "isotopeNameNotDefined"
		mat.Specs = compound

		check(t, testCase{
			Input: createMaterialMap(mat),
			ExpectedError: errors.New(
				"[serializer] Material{Id: 1} -> mat.dat: \"isotopeNameNotDefined\"" +
					" isotope mapping to shield format not found",
			),
		})
	})

	t.Run("ExternalStoppingPowerFromPredefinedMaterialMappingNotFound", func(t *testing.T) {
		const id = 1
		mat := genCompound(id)
		compound := mat.Specs.(specs.MaterialCompound)
		mat.Specs = compound
		check(t, testCase{
			Input: createMaterialMap(mat),
			ExpectedError: errors.New(
				"[serializer] Material{Id: 1} -> mat.dat: \"espfpNameNotDefined\"" +
					" material mapping to shield format not found",
			),
		})
	})

}

func genSimplePredefined(id int64) specs.Material {
	return specs.Material{
		ID: specs.MaterialID(id),
		Specs: specs.MaterialPredefined{
			PredefinedID: "urea",
		},
	}
}

func genFullPredefined(id int64) specs.Material {
	return specs.Material{
		ID: specs.MaterialID(id),
		Specs: specs.MaterialPredefined{
			PredefinedID:  "methanol",
			StateOfMatter: specs.Liquid,
			Density:       func() *float64 { a := 123.45; return &a }(),
		},
	}
}

func genCompound(id int64) specs.Material {
	return specs.Material{
		ID: specs.MaterialID(id),
		Specs: specs.MaterialCompound{
			Density:       99.9,
			StateOfMatter: specs.Solid,
			Elements: []specs.Element{
				specs.Element{
					Isotope: "gd-*", RelativeStoichiometricFraction: 2,
					AtomicMass: func() *int64 { a := int64(10); return &a }(),
				},
				specs.Element{
					Isotope: "u-235", RelativeStoichiometricFraction: 123,
					IValue: func() *float64 { a := float64(555.34); return &a }(),
				},
			},
		},
	}
}

func genAnotherCompound(id int64) specs.Material {
	return specs.Material{
		ID: specs.MaterialID(id),
		Specs: specs.MaterialCompound{
			Name:          "pies",
			Density:       0.999,
			StateOfMatter: specs.Gas,
			Elements: []specs.Element{
				specs.Element{
					Isotope: "c-*",
					RelativeStoichiometricFraction: 4,
					AtomicMass:                     func() *int64 { a := int64(1); return &a }(),
				},
				specs.Element{
					Isotope: "si-*",
					RelativeStoichiometricFraction: 1,
					IValue: func() *float64 { a := float64(0.34); return &a }(),
				},
				specs.Element{
					Isotope: "na-23",
					RelativeStoichiometricFraction: 11111,
					AtomicMass:                     func() *int64 { a := int64(3); return &a }(),
					IValue:                         func() *float64 { a := float64(987.654); return &a }(),
				},
			},
		},
	}
}

func genVacuum(id int64) specs.Material {
	return specs.Material{
		ID: specs.MaterialID(id),
		Specs: specs.MaterialPredefined{
			PredefinedID: "vacuum",
		},
	}
}

// TODO(simple): remove
func createMaterialMap(materials ...specs.Material) map[specs.MaterialID]specs.Material {
	res := map[specs.MaterialID]specs.Material{}
	for _, m := range materials {
		res[m.ID] = m
	}
	return res
}

func setPredefinedID(mat PredefinedMaterial, id ShieldID) PredefinedMaterial {
	mat.ID = id
	return mat
}

func setCompoundID(mat CompoundMaterial, id ShieldID) CompoundMaterial {
	mat.ID = id
	return mat
}
