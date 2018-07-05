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
		m("id", Int64Marshaler(int64(z.ID)))
		m("parentId", Int64Marshaler(int64(z.ParentID)))
		m("baseId", Int64Marshaler(int64(z.BaseID)))
		m("materialId", Int64Marshaler(int64(z.MaterialID)))
		m("construction",
			ListMarshaler(z.Construction, func(z setup.ZoneOperation) marshaler {
				return StructMarshaler(func(m fieldMarshaler) {
					m("bodyId", Int64Marshaler(int64(z.BodyID)))
					m("type", EnumMarshaler(z.Type, mapOperationToJSON))
				})
			}),
		)
	})
}

func ZoneUnmarshaler(z *setup.Zone) unmarshaler {
	return StructUnmarshaler(func(u fieldUnmarshaler) {
		u("id", Int64Unmarshaler((*int64)(&z.ID)))
		u("parentId", Int64Unmarshaler((*int64)(&z.ParentID)))
		u("baseId", Int64Unmarshaler((*int64)(&z.BaseID)))
		u("materialId", Int64Unmarshaler((*int64)(&z.MaterialID)))
		u("construction",
			ListUnmarshaler(&z.Construction, func(z *setup.ZoneOperation) unmarshaler {
				return StructUnmarshaler(func(u fieldUnmarshaler) {
					u("bodyId", Int64Unmarshaler((*int64)(&z.BodyID)))
					u("type", EnumUnmarshaler(&z.Type, mapOperationToJSON))
				})
			}),
		)
	})
}
