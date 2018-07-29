package simulation

import (
	"github.com/yaptide/yaptide/pkg/converter/specs"
	. "github.com/yaptide/yaptide/pkg/serialize"
)

var mapOperationToJSON = map[specs.ZoneOperationType]string{
	specs.Intersect: "intersect",
	specs.Subtract:  "subtract",
	specs.Union:     "union",
}

var mapJSONToOperation = map[string]specs.ZoneOperationType{
	"intersect": specs.Intersect,
	"subtract":  specs.Subtract,
	"union":     specs.Union,
}

func zoneMarshaler(z Zone) marshaler {
	return StructMarshaler(func(m fieldMarshaler) {
		m("id", Int64Marshaler(int64(z.ID)))
		m("name", StringMarshaler(z.Name))
		m("parentId", Int64Marshaler(int64(z.ParentID)))
		m("baseId", Int64Marshaler(int64(z.BaseID)))
		m("materialId", Int64Marshaler(int64(z.MaterialID)))
		m("construction",
			ListMarshaler(z.Construction, func(z specs.ZoneOperation) marshaler {
				return StructMarshaler(func(m fieldMarshaler) {
					m("bodyId", Int64Marshaler(int64(z.BodyID)))
					m("type", EnumMarshaler(z.Type, mapOperationToJSON))
				})
			}),
		)
	})
}

func zoneUnmarshaler(z *Zone) unmarshaler {
	return StructUnmarshaler(func(u fieldUnmarshaler) {
		u("id", Int64Unmarshaler((*int64)(&z.ID)))
		u("name", StringUnmarshaler(&z.Name))
		u("parentId", Int64Unmarshaler((*int64)(&z.ParentID)))
		u("baseId", Int64Unmarshaler((*int64)(&z.BaseID)))
		u("materialId", Int64Unmarshaler((*int64)(&z.MaterialID)))
		u("construction",
			ListUnmarshaler(&z.Construction, func(z *specs.ZoneOperation) unmarshaler {
				return StructUnmarshaler(func(u fieldUnmarshaler) {
					u("bodyId", Int64Unmarshaler((*int64)(&z.BodyID)))
					u("type", EnumUnmarshaler(&z.Type, mapOperationToJSON))
				})
			}),
		)
	})
}
