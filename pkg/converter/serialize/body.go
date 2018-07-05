package serialize

import (
	"fmt"

	"github.com/yaptide/yaptide/pkg/converter/setup"
)

var bodyType = struct {
	cuboid   string
	cylinder string
	sphere   string
}{
	cuboid:   "cuboid",
	cylinder: "cylinder",
	sphere:   "sphere",
}

func BodyMarshaler(b setup.Body) marshaler {
	return StructMarshaler(func(m fieldMarshaler) {
		m("id", Int64Marshaler(int64(b.ID)))
		m("geometry", func() (interface{}, error) {
			switch b := b.Geometry.(type) {
			case setup.SphereBody:
				return BodySphereMarshaler(b)()
			case setup.CuboidBody:
				return BodyCuboidMarshaler(b)()
			case setup.CylinderBody:
				return BodyCylinderMarshaler(b)()
			default:
				return nil, fmt.Errorf("unknown body geometry")
			}
		})
	})
}

func BodyUnmarshaler(b *setup.Body) unmarshaler {
	return StructUnmarshaler(func(u fieldUnmarshaler) {
		u("id", Int64Unmarshaler((*int64)(&b.ID)))
		u("geometry", UnionTypeUnmarshaler(
			func(unionType string) unmarshaler {
				switch unionType {
				case bodyType.sphere:
					return BodySphereUnmarshaler(&b.Geometry)
				case bodyType.cuboid:
					return BodyCuboidUnmarshaler(&b.Geometry)
				case bodyType.cylinder:
					return BodyCylinderUnmarshaler(&b.Geometry)
				default:
					return func(raw interface{}) error {
						return fmt.Errorf("unknown body type")
					}
				}
			}))
	})
}

func BodySphereMarshaler(b setup.SphereBody) marshaler {
	return StructMarshaler(func(m fieldMarshaler) {
		m("radius", Float64Marshaler(b.Radius))
		m("center", PointMarshaler(b.Center))
	})
}

func BodySphereUnmarshaler(b *setup.BodyGeometry) unmarshaler {
	*b = setup.SphereBody{}
	return StructUnmarshaler(func(u fieldUnmarshaler) {
		b := (*b).(setup.SphereBody)
		u("radius", Float64Unmarshaler(&b.Radius))
		u("center", PointUnmarshaler(&b.Center))
	})
}

func BodyCuboidMarshaler(b setup.CuboidBody) marshaler {
	return StructMarshaler(func(m fieldMarshaler) {
		m("center", PointMarshaler(b.Center))
		m("size", Vec3DMarshaler(b.Size))
	})
}

func BodyCuboidUnmarshaler(b *setup.BodyGeometry) unmarshaler {
	*b = setup.CuboidBody{}
	return StructUnmarshaler(func(u fieldUnmarshaler) {
		b := (*b).(setup.CuboidBody)
		u("center", PointUnmarshaler(&b.Center))
		u("size", Vec3DUnmarshaler(&b.Size))
	})
}

func BodyCylinderMarshaler(b setup.CylinderBody) marshaler {
	return StructMarshaler(func(m fieldMarshaler) {
		m("height", Float64Marshaler(b.Height))
		m("radius", Float64Marshaler(b.Radius))
		m("center", PointMarshaler(b.Center))
	})
}

func BodyCylinderUnmarshaler(b *setup.BodyGeometry) unmarshaler {
	*b = setup.CylinderBody{}
	return StructUnmarshaler(func(u fieldUnmarshaler) {
		b := (*b).(setup.CylinderBody)
		u("height", Float64Unmarshaler(&b.Height))
		u("radius", Float64Unmarshaler(&b.Radius))
		u("center", PointUnmarshaler(&b.Center))
	})
}
