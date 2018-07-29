package material

import (
	"fmt"
	"sort"

	"github.com/yaptide/yaptide/pkg/converter/shield/mapping"
	"github.com/yaptide/yaptide/pkg/converter/specs"
)

// ShieldID ...
type ShieldID int

// Materials contains representation of specs.MaterialsMap,
// which is easily serializable in shield serializer.
type Materials struct {
	Predefined []PredefinedMaterial
	Compound   []CompoundMaterial
}

// PredefinedMaterial represent specs.Predefined.
type PredefinedMaterial struct {
	ID            ShieldID
	ICRUNumber    mapping.MaterialICRU
	StateOfMatter mapping.StateOfMatter
	Density       *float64
}

// Element represent specs.Element.
type Element struct {
	ID                             mapping.IsotopeNUCLID
	RelativeStoichiometricFraction int64
	AtomicMass                     *int64
	IValue                         *float64
}

// CompoundMaterial represent specs.Compound.
type CompoundMaterial struct {
	ID            ShieldID
	StateOfMatter mapping.StateOfMatter
	Density       float64
	Elements      []Element
}

// ConvertMaterials ...
func ConvertMaterials(
	specsMat map[specs.MaterialID]specs.Material,
) (Materials, map[specs.MaterialID]ShieldID, error) {
	result := Materials{
		Predefined: []PredefinedMaterial{},
		Compound:   []CompoundMaterial{},
	}

	materialIDToShield := map[specs.MaterialID]ShieldID{}

	const maxMaterialsNumber = 100

	if len(specsMat) > maxMaterialsNumber {
		return Materials{}, nil, fmt.Errorf(
			"Only %d distinct materials are permitted in shield (%d > %d)",
			maxMaterialsNumber,
			len(specsMat),
			maxMaterialsNumber,
		)
	}

	predefMaterialsIds := []specs.MaterialID{}
	compoundMaterialsIds := []specs.MaterialID{}

	for id, mat := range specsMat {
		switch g := mat.Specs.(type) {
		case specs.MaterialPredefined:
			if g.PredefinedID != "vacuum" {
				predefMaterialsIds = append(predefMaterialsIds, id)
			} else {
				materialIDToShield[id] = ShieldID(mapping.PredefinedMaterialsToShieldICRU["vacuum"])
			}
		case specs.MaterialCompound:
			compoundMaterialsIds = append(compoundMaterialsIds, id)
		case specs.MaterialVoxel:
			return Materials{}, nil, fmt.Errorf(
				"Voxel material serialization not implemented",
			)
		default:
			return Materials{}, nil, fmt.Errorf(
				"Unknown material type",
			)
		}
	}

	nextShieldID := 1
	for _, ids := range [][]specs.MaterialID{predefMaterialsIds, compoundMaterialsIds} {
		sort.SliceStable(ids, func(i, j int) bool { return ids[i] < ids[j] })
		for _, id := range ids {
			materialIDToShield[id] = ShieldID(nextShieldID)
			nextShieldID++
		}
	}

	for _, predefID := range predefMaterialsIds {
		predef, err := createPredefinedMaterial(
			specsMat[predefID].Specs.(specs.MaterialPredefined), materialIDToShield[predefID],
		)
		if err != nil {
			return Materials{}, nil, err
		}
		result.Predefined = append(result.Predefined, predef)
	}

	for _, compoundID := range compoundMaterialsIds {
		compound, err := createCompoundMaterial(
			specsMat[compoundID].Specs.(specs.MaterialCompound), materialIDToShield[compoundID],
		)
		if err != nil {
			return Materials{}, nil, err
		}
		result.Compound = append(result.Compound, compound)
	}
	return result, materialIDToShield, nil
}

// SerializeStateOfMatter return true, if StateOfMatter should be serialized.
func (p *PredefinedMaterial) SerializeStateOfMatter() bool {
	return p.StateOfMatter != mapping.StateOfMatterToShield[specs.UndefinedStateOfMatter]
}

// SerializeDensity return true, if Density should be serialized.
func (p *PredefinedMaterial) SerializeDensity() bool {
	return *p.Density > 0.0
}

// SerializeAtomicMass return true, if AtomicMass should be serialized.
func (e *Element) SerializeAtomicMass() bool {
	return e.AtomicMass != nil
}

// SerializeIValue return true, if IValue should be serialized.
func (e *Element) SerializeIValue() bool {
	return e.IValue != nil
}

func createPredefinedMaterial(
	predef specs.MaterialPredefined, id ShieldID,
) (PredefinedMaterial, error) {
	ICRUNumber, found := mapping.PredefinedMaterialsToShieldICRU[predef.PredefinedID]
	if !found {
		return PredefinedMaterial{}, fmt.Errorf(
			"\"%s\" material mapping to shield format not found", predef.PredefinedID,
		)
	}

	return PredefinedMaterial{
		ID:            id,
		StateOfMatter: mapping.StateOfMatterToShield[predef.StateOfMatter],
		Density:       predef.Density,
		ICRUNumber:    ICRUNumber,
	}, nil
}

func createCompoundMaterial(
	compound specs.MaterialCompound, id ShieldID,
) (CompoundMaterial, error) {
	const maxElementsNumber = 13

	if compound.StateOfMatter == specs.UndefinedStateOfMatter {
		return CompoundMaterial{}, fmt.Errorf(
			"StateOfMatter must be defined for Compound material",
		)
	}
	if compound.Density <= 0.0 {
		return CompoundMaterial{}, fmt.Errorf(
			"Density must be specified for Compund material",
		)
	}
	if len(compound.Elements) > maxElementsNumber {
		return CompoundMaterial{}, fmt.Errorf(
			"Only %d elements for Compound are permitted in shield (%d > %d)",
			maxElementsNumber, len(compound.Elements), maxElementsNumber,
		)
	}

	elements := []Element{}
	for _, element := range compound.Elements {
		isotopeNUCLID, found := mapping.IsotopesToShieldNUCLID[element.Isotope]
		if !found {
			return CompoundMaterial{}, fmt.Errorf(
				"\"%s\" isotope mapping to shield format not found", element.Isotope,
			)
		}
		elements = append(elements, Element{
			ID: isotopeNUCLID,
			RelativeStoichiometricFraction: element.RelativeStoichiometricFraction,
			AtomicMass:                     element.AtomicMass,
			IValue:                         element.IValue,
		})
	}

	return CompoundMaterial{
		ID:            id,
		StateOfMatter: mapping.StateOfMatterToShield[compound.StateOfMatter],
		Density:       compound.Density,
		Elements:      elements,
	}, nil
}
