package simulation

import (
	"fmt"

	"github.com/yaptide/yaptide/pkg/converter/specs"
	. "github.com/yaptide/yaptide/pkg/serialize"
)

var materialType = struct {
	predefined string
	compound   string
	voxel      string
}{
	predefined: "predefined",
	compound:   "compound",
	voxel:      "voxel",
}

func materialMarshaler(mat Material) marshaler {
	return StructMarshaler(func(m fieldMarshaler) {
		m("id", Int64Marshaler(int64(mat.ID)))
		m("name", StringMarshaler(mat.Name))
		m("specs", func() (interface{}, error) {
			switch mat := mat.Specs.(type) {
			case specs.MaterialPredefined:
				return materialPredfinedMarshaler(mat)()
			case specs.MaterialCompound:
				return materialCompoundMarshaler(mat)()
			case specs.MaterialVoxel:
				return materialVoxelMarshaler(mat)()
			default:
				return nil, fmt.Errorf("unknown material %T", mat)
			}
		})
	})
}

func materialUnmarshaler(m *Material) unmarshaler {
	return StructUnmarshaler(func(u fieldUnmarshaler) {
		u("id", Int64Unmarshaler((*int64)(&m.ID)))
		u("name", StringUnmarshaler(&m.Name))
		u("specs", UnionTypeUnmarshaler(func(unionType string) unmarshaler {
			switch unionType {
			case materialType.predefined:
				return materialPredfinedUnmarshaler(&m.Specs)
			case materialType.compound:
				return materialCompoundUnmarshaler(&m.Specs)
			case materialType.voxel:
				return materialVoxelUnmarshaler(&m.Specs)
			default:
				return func(raw interface{}) error {
					return fmt.Errorf("unknown material type")
				}
			}
		}))
	})
}

func materialPredfinedMarshaler(o specs.MaterialPredefined) marshaler {
	return StructMarshaler(func(m fieldMarshaler) {
		m("type", StringMarshaler(materialType.predefined))
		m("predefinedId", StringMarshaler(o.PredefinedID))
		m("density", PtrMarshaler(o.Density, Float64Marshaler))
		m("stateOfMatter", stateOfMatterMarshaler(o.StateOfMatter))
	})
}

func materialPredfinedUnmarshaler(o *specs.MaterialSpecs) unmarshaler {
	*o = specs.MaterialPredefined{}
	return StructUnmarshaler(func(u fieldUnmarshaler) {
		o := (*o).(specs.MaterialPredefined)
		u("predefinedId", StringUnmarshaler(&o.PredefinedID))
		u("density", PtrUnmarshaler(&o.Density, Float64Unmarshaler))
		u("stateOfMatter", stateOfMatterUnmarshaler(&o.StateOfMatter))
	})
}

func materialCompoundMarshaler(o specs.MaterialCompound) marshaler {
	return StructMarshaler(func(m fieldMarshaler) {
		m("type", StringMarshaler(materialType.compound))
		m("density", Float64Marshaler(o.Density))
		m("stateOfMatter", stateOfMatterMarshaler(o.StateOfMatter))
		m("elements",
			ListMarshaler(o.Elements, func(o specs.Element) marshaler {
				return StructMarshaler(func(m fieldMarshaler) {
					m("isotope", StringMarshaler(o.Isotope))
					m(
						"relativeStoichiometricFraction",
						Int64Marshaler(o.RelativeStoichiometricFraction),
					)
					m("atomicMass", PtrMarshaler(o.AtomicMass, Int64Marshaler))
					m("iValue", PtrMarshaler(o.IValue, Float64Marshaler))
				})
			}),
		)
	})
}

func materialCompoundUnmarshaler(o *specs.MaterialSpecs) unmarshaler {
	*o = specs.MaterialCompound{}
	return StructUnmarshaler(func(u fieldUnmarshaler) {
		o := (*o).(specs.MaterialCompound)
		u("name", StringUnmarshaler(&o.Name))
		u("density", Float64Unmarshaler(&o.Density))
		u("stateOfMatter", stateOfMatterUnmarshaler(&o.StateOfMatter))
		u("elements",
			ListUnmarshaler(&o.Elements, func(o *specs.Element) unmarshaler {
				return StructUnmarshaler(func(u fieldUnmarshaler) {
					u("isotope", StringUnmarshaler(&o.Isotope))
					u(
						"relativeStoichiometricFraction",
						Int64Unmarshaler(&o.RelativeStoichiometricFraction),
					)
					u("atomicMass", PtrUnmarshaler(&o.AtomicMass, Int64Unmarshaler))
					u("iValue", PtrUnmarshaler(&o.IValue, Float64Unmarshaler))
				})
			}),
		)
	})
}

func materialVoxelMarshaler(o specs.MaterialVoxel) marshaler {
	return func() (interface{}, error) {
		return "", fmt.Errorf("unsuported type voxel")
	}
}

func materialVoxelUnmarshaler(o *specs.MaterialSpecs) unmarshaler {
	return func(raw interface{}) error {
		return fmt.Errorf("unsuported type voxel")
	}
}
