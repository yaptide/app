package serialize

import (
	"fmt"

	"github.com/yaptide/yaptide/pkg/converter/setup"
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

func MaterialMarshaler(mat setup.Material) marshaler {
	return StructMarshaler(func(m fieldMarshaler) {
		m("id", Int64Marshaler(int64(mat.ID)))
		m("specs", func() (interface{}, error) {
			switch mat := mat.Specs.(type) {
			case setup.MaterialPredefined:
				return MaterialPredfinedMarshaler(mat)()
			case setup.MaterialCompound:
				return MaterialCompoundMarshaler(mat)()
			case setup.MaterialVoxel:
				return MaterialVoxelMarshaler(mat)()
			default:
				return nil, fmt.Errorf("unknown material")
			}
		})
	})
}

func MaterialUnmarshaler(m *setup.Material) unmarshaler {
	return StructUnmarshaler(func(u fieldUnmarshaler) {
		u("id", Int64Unmarshaler((*int64)(&m.ID)))
		u("specs", UnionTypeUnmarshaler(func(unionType string) unmarshaler {
			switch unionType {
			case materialType.predefined:
				return MaterialPredfinedUnmarshaler(&m.Specs)
			case materialType.compound:
				return MaterialCompoundUnmarshaler(&m.Specs)
			case materialType.voxel:
				return MaterialVoxelUnmarshaler(&m.Specs)
			default:
				return func(raw interface{}) error {
					return fmt.Errorf("unknown material type")
				}
			}
		}))
	})
}

func MaterialPredfinedMarshaler(o setup.MaterialPredefined) marshaler {
	return StructMarshaler(func(m fieldMarshaler) {
		m("type", StringMarshaler(materialType.predefined))
		m("predefinedId", StringMarshaler(o.PredefinedID))
		m("density", Float64Marshaler(o.Density))
		m("stateOfMatter", StateOfMatterMarshaler(o.StateOfMatter))
		m("loadExternalStoppingPower", BooleanMarshaler(o.LoadExternalStoppingPower))
	})
}

func MaterialPredfinedUnmarshaler(o *setup.MaterialSpecs) unmarshaler {
	*o = setup.MaterialPredefined{}
	return StructUnmarshaler(func(u fieldUnmarshaler) {
		o := (*o).(setup.MaterialPredefined)
		u("predefinedId", StringUnmarshaler(&o.PredefinedID))
		u("density", Float64Unmarshaler(&o.Density))
		u("stateOfMatter", StateOfMatterUnmarshaler(&o.StateOfMatter))
		u("loadExternalStoppingPower", BooleanUnmarshaler(&o.LoadExternalStoppingPower))
	})
}

func MaterialCompoundMarshaler(o setup.MaterialCompound) marshaler {
	return StructMarshaler(func(m fieldMarshaler) {
		m("type", StringMarshaler(materialType.compound))
		m("name", StringMarshaler(o.Name))
		m("density", Float64Marshaler(o.Density))
		m("stateOfMatter", StateOfMatterMarshaler(o.StateOfMatter))
		m(
			"externalStopingPowerFromPredefined",
			StringMarshaler(o.ExternalStoppingPowerFromPredefined),
		)
		m("elements",
			ListMarshaler(o.Elements, func(o setup.Element) marshaler {
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

func MaterialCompoundUnmarshaler(o *setup.MaterialSpecs) unmarshaler {
	*o = setup.MaterialCompound{}
	return StructUnmarshaler(func(u fieldUnmarshaler) {
		o := (*o).(setup.MaterialCompound)
		u("name", StringUnmarshaler(&o.Name))
		u("density", Float64Unmarshaler(&o.Density))
		u("stateOfMatter", StateOfMatterUnmarshaler(&o.StateOfMatter))
		u(
			"externalStopingPowerFromPredefined",
			StringUnmarshaler(&o.ExternalStoppingPowerFromPredefined),
		)
		u("elements",
			ListUnmarshaler(&o.Elements, func(o *setup.Element) unmarshaler {
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

func MaterialVoxelMarshaler(o setup.MaterialVoxel) marshaler {
	return func() (interface{}, error) {
		return "", fmt.Errorf("unsuported type voxel")
	}
}

func MaterialVoxelUnmarshaler(o *setup.MaterialSpecs) unmarshaler {
	return func(raw interface{}) error {
		return fmt.Errorf("unsuported type voxel")
	}
}

var mapStateOfMatterToJSON = map[setup.StateOfMatter]string{
	setup.Gas:    "gas",
	setup.Liquid: "liquid",
	setup.Solid:  "solid",
}

var mapJSONToStateOfMatter = map[string]setup.StateOfMatter{
	"gas":    setup.Gas,
	"liquid": setup.Liquid,
	"solid":  setup.Solid,
}

func StateOfMatterMarshaler(o setup.StateOfMatter) marshaler {
	return EnumMarshaler(o, mapStateOfMatterToJSON)
}

func StateOfMatterUnmarshaler(o *setup.StateOfMatter) unmarshaler {
	return EnumUnmarshaler(o, mapJSONToStateOfMatter)
}
