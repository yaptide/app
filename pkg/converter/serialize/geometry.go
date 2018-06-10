package serialize

import "github.com/yaptide/yaptide/pkg/converter/geometry"

func PointMarshaler(p geometry.Point) marshaler {
	return StructMarshaler(func(m fieldMarshaler) {
		m("x", Float64Marshaler(p.X))
		m("y", Float64Marshaler(p.Y))
		m("z", Float64Marshaler(p.Z))
	})
}

func PointUnmarshaler(p *geometry.Point) unmarshaler {
	return StructUnmarshaler(func(u fieldUnmarshaler) {
		u("x", Float64Unmarshaler(&p.X))
		u("y", Float64Unmarshaler(&p.Y))
		u("z", Float64Unmarshaler(&p.Z))
	})
}

func Vec3DMarshaler(p geometry.Vec3D) marshaler {
	return StructMarshaler(func(m fieldMarshaler) {
		m("x", Float64Marshaler(p.X))
		m("y", Float64Marshaler(p.Y))
		m("z", Float64Marshaler(p.Z))
	})
}

func Vec3DUnmarshaler(p *geometry.Vec3D) unmarshaler {
	return StructUnmarshaler(func(u fieldUnmarshaler) {
		u("x", Float64Unmarshaler(&p.X))
		u("y", Float64Unmarshaler(&p.Y))
		u("z", Float64Unmarshaler(&p.Z))
	})
}

func Vec3DIntMarshaler(p geometry.Vec3DInt) marshaler {
	return StructMarshaler(func(m fieldMarshaler) {
		m("x", Int64Marshaler(p.X))
		m("y", Int64Marshaler(p.Y))
		m("z", Int64Marshaler(p.Z))
	})
}

func Vec3DIntUnmarshaler(p *geometry.Vec3DInt) unmarshaler {
	return StructUnmarshaler(func(u fieldUnmarshaler) {
		u("x", Int64Unmarshaler(&p.X))
		u("y", Int64Unmarshaler(&p.Y))
		u("z", Int64Unmarshaler(&p.Z))
	})
}

func Vec3DCylindricalIntMarshaler(p geometry.Vec3DCylindricalInt) marshaler {
	return StructMarshaler(func(m fieldMarshaler) {
		m("radius", Int64Marshaler(p.Radius))
		m("angle", Int64Marshaler(p.Angle))
		m("z", Int64Marshaler(p.Z))
	})
}

func Vec3DCylindricalIntUnmarshaler(p *geometry.Vec3DCylindricalInt) unmarshaler {
	return StructUnmarshaler(func(u fieldUnmarshaler) {
		u("radius", Int64Unmarshaler(&p.Radius))
		u("angle", Int64Unmarshaler(&p.Angle))
		u("z", Int64Unmarshaler(&p.Z))
	})
}

func RangeMarshaler(p geometry.Range) marshaler {
	return StructMarshaler(func(m fieldMarshaler) {
		m("min", Float64Marshaler(p.Min))
		m("max", Float64Marshaler(p.Max))
	})
}

func RangeUnmarshaler(p *geometry.Range) unmarshaler {
	return StructUnmarshaler(func(u fieldUnmarshaler) {
		u("min", Float64Unmarshaler(&p.Min))
		u("max", Float64Unmarshaler(&p.Max))
	})
}
