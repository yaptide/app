package simulation

import (
	. "github.com/yaptide/yaptide/pkg/serialize"
)

// Color represent (R, G, B , A) color.
type Color struct {
	R int64
	G int64
	B int64
	A int64
}

// NewColor construct new color.
func NewColor(R, G, B, A int64) Color {
	return Color{R, G, B, A}
}

var (
	// White color.
	White = NewColor(0xFF, 0xFF, 0xFF, 0xFF)

	//Gray color.
	Gray = NewColor(0x80, 0x80, 0x80, 0xFF)
)

func colorMarshaler(c Color) marshaler {
	return StructMarshaler(func(m fieldMarshaler) {
		m("r", Int64Marshaler(c.R))
		m("g", Int64Marshaler(c.G))
		m("b", Int64Marshaler(c.B))
		m("a", Int64Marshaler(c.A))
	})
}

func colorUnmarshaler(c *Color) unmarshaler {
	return StructUnmarshaler(func(u fieldUnmarshaler) {
		u("r", Int64Unmarshaler(&c.R))
		u("g", Int64Unmarshaler(&c.G))
		u("b", Int64Unmarshaler(&c.B))
		u("a", Int64Unmarshaler(&c.A))
	})
}
