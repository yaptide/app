package serialize

import (
	"github.com/yaptide/yaptide/pkg/converter/setup"
)

func MarshalSetup(s setup.Setup) (interface{}, error) {
	return StructMarshaler(func(m fieldMarshaler) {
		m("materials", ListMarshaler(s.Materials, MaterialMarshaler))
		m("zones", ListMarshaler(s.Zones, ZoneMarshaler))
		m("bodies", ListMarshaler(s.Bodies, BodyMarshaler))
		m("detectors", ListMarshaler(s.Detectors, DetectorMarshaler))
		m("beam", BeamMarshaler(s.Beam))
		m("options", OptionsMarshaler(s.Options))
	})()
}

func UnmarshalSetup(raw interface{}, s *setup.Setup) error {
	return StructUnmarshaler(func(u fieldUnmarshaler) {
		u("materials", ListUnmarshaler(&s.Materials, MaterialUnmarshaler))
		u("zones", ListUnmarshaler(&s.Zones, ZoneUnmarshaler))
		u("bodies", ListUnmarshaler(&s.Bodies, BodyUnmarshaler))
		u("detectors", ListUnmarshaler(&s.Detectors, DetectorUnmarshaler))
		u("beam", BeamUnmarshaler(&s.Beam))
		u("options", OptionsUnmarshaler(&s.Options))
	})(raw)
}
