package specs

import (
	"github.com/yaptide/yaptide/pkg/converter/geometry"
)

// DetectorPlane detector.
type DetectorPlane struct {
	Point  geometry.Point
	Normal geometry.Vec3D
}

// MarshalJSON json.Marshaller implementation.
func (d DetectorPlane) Validate() error {
	return nil
}
