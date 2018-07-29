package simulation

import (
	. "github.com/yaptide/yaptide/pkg/serialize"
)

func specsMarshaler(s Specs) marshaler {
	return StructMarshaler(func(m fieldMarshaler) {
		m("materials", ListMarshaler(s.Materials, materialMarshaler))
		m("zones", ListMarshaler(s.Zones, zoneMarshaler))
		m("bodies", ListMarshaler(s.Bodies, bodyMarshaler))
		m("detectors", ListMarshaler(s.Detectors, detectorMarshaler))
		m("beam", beamMarshaler(s.Beam))
		m("options", optionsMarshaler(s.Options))
	})
}

func specsUnmarshaler(s *Specs) unmarshaler {
	return StructUnmarshaler(func(u fieldUnmarshaler) {
		u("materials", ListUnmarshaler(&s.Materials, materialUnmarshaler))
		u("zones", ListUnmarshaler(&s.Zones, zoneUnmarshaler))
		u("bodies", ListUnmarshaler(&s.Bodies, bodyUnmarshaler))
		u("detectors", ListUnmarshaler(&s.Detectors, detectorUnmarshaler))
		u("beam", beamUnmarshaler(&s.Beam))
		u("options", optionsUnmarshaler(&s.Options))
	})
}
