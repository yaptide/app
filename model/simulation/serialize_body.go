package simulation

import (
	"fmt"

	"github.com/yaptide/yaptide/pkg/converter/specs"
	. "github.com/yaptide/yaptide/pkg/serialize"
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

func bodyMarshaler(b specs.Body) marshaler {
	return StructMarshaler(func(m fieldMarshaler) {
		m("id", Int64Marshaler(int64(b.ID)))
		m("geometry", func() (interface{}, error) {
			switch b := b.Geometry.(type) {
			case specs.BodySphere:
				return bodySphereMarshaler(b)()
			case specs.BodyCuboid:
				return bodyCuboidMarshaler(b)()
			case specs.BodyCylinder:
				return bodyCylinderMarshaler(b)()
			default:
				return nil, fmt.Errorf("unknown body geometry")
			}
		})
	})
}

func bodyUnmarshaler(b *specs.Body) unmarshaler {
	return StructUnmarshaler(func(u fieldUnmarshaler) {
		u("id", Int64Unmarshaler((*int64)(&b.ID)))
		u("geometry", UnionTypeUnmarshaler(
			func(unionType string) unmarshaler {
				switch unionType {
				case bodyType.sphere:
					return bodySphereUnmarshaler(&b.Geometry)
				case bodyType.cuboid:
					return bodyCuboidUnmarshaler(&b.Geometry)
				case bodyType.cylinder:
					return bodyCylinderUnmarshaler(&b.Geometry)
				default:
					return func(raw interface{}) error {
						return fmt.Errorf("unknown body type")
					}
				}
			}))
	})
}

func bodySphereMarshaler(b specs.BodySphere) marshaler {
	return StructMarshaler(func(m fieldMarshaler) {
		m("type", StringMarshaler(bodyType.sphere))
		m("radius", Float64Marshaler(b.Radius))
		m("center", pointMarshaler(b.Center))
	})
}

func bodySphereUnmarshaler(body *specs.BodyGeometry) unmarshaler {
	return StructUnmarshaler(func(u fieldUnmarshaler) {
		b := specs.BodySphere{}
		u("radius", Float64Unmarshaler(&b.Radius))
		u("center", pointUnmarshaler(&b.Center))
		*body = b
	})
}

func bodyCuboidMarshaler(b specs.BodyCuboid) marshaler {
	return StructMarshaler(func(m fieldMarshaler) {
		m("type", StringMarshaler(bodyType.cuboid))
		m("center", pointMarshaler(b.Center))
		m("size", vec3DMarshaler(b.Size))
	})
}

func bodyCuboidUnmarshaler(body *specs.BodyGeometry) unmarshaler {
	return StructUnmarshaler(func(u fieldUnmarshaler) {
		b := specs.BodyCuboid{}
		u("center", pointUnmarshaler(&b.Center))
		u("size", vec3DUnmarshaler(&b.Size))
		*body = b
	})
}

func bodyCylinderMarshaler(b specs.BodyCylinder) marshaler {
	return StructMarshaler(func(m fieldMarshaler) {
		m("type", StringMarshaler(bodyType.cylinder))
		m("height", Float64Marshaler(b.Height))
		m("radius", Float64Marshaler(b.Radius))
		m("center", pointMarshaler(b.Center))
	})
}

func bodyCylinderUnmarshaler(body *specs.BodyGeometry) unmarshaler {
	return StructUnmarshaler(func(u fieldUnmarshaler) {
		b := specs.BodyCylinder{}
		u("height", Float64Unmarshaler(&b.Height))
		u("radius", Float64Unmarshaler(&b.Radius))
		u("center", pointUnmarshaler(&b.Center))
		*body = b
	})
}
