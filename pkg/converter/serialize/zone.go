package serialize

import "github.com/yaptide/yaptide/pkg/converter/setup"

var mapOperationToJSON = map[setup.ZoneOperationType]string{
	setup.Intersect: "intersect",
	setup.Subtract:  "subtract",
	setup.Union:     "union",
}

var mapJSONToOperation = map[string]setup.ZoneOperationType{
	"intersect": setup.Intersect,
	"subtract":  setup.Subtract,
	"union":     setup.Union,
}

func ZoneMarshaler(z setup.Zone) marshaler {
	return StructMarshaler(func(m fieldMarshaler) {
		m("id", Int64Marshaler(z.ID))
		m("parentId", Int64Marshaler(z.ParentID))
		m("baseId", Int64Marshaler(z.BaseID))
		m("materialId", Int64Marshaler(z.MaterialID))
		m("construction",
			ListMarshaler(z.Construction, func(z setup.ZoneOperation) marshaler {
				return StructMarshaler(func(m fieldMarshaler) {
					m("bodyId", Int64Marshaler(z.BodyID))
					m("type", EnumMarshaler(z.Type, mapOperationToJSON))
				})
			}),
		)
	})
}

func ZoneUnmarshaler(z *setup.Zone) unmarshaler {
	return StructUnmarshaler(func(u fieldUnmarshaler) {
		u("id", Int64Unmarshaler(&z.ID))
		u("parentId", Int64Unmarshaler(&z.ParentID))
		u("baseId", Int64Unmarshaler(&z.BaseID))
		u("materialId", Int64Unmarshaler(&z.MaterialID))
		u("construction",
			ListUnmarshaler(&z.Construction, func(z *setup.ZoneOperation) unmarshaler {
				return StructUnmarshaler(func(u fieldUnmarshaler) {
					u("bodyId", Int64Unmarshaler(&z.BodyID))
					u("type", EnumUnmarshaler(&z.Type, mapOperationToJSON))
				})
			}),
		)
	})
}
