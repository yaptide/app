package setup

import (
	"github.com/yaptide/yaptide/pkg/converter/geometry"
)

// DetectorPlane detector.
type DetectorPlane struct {
	Point  geometry.Point `json:"point"`
	Normal geometry.Vec3D `json:"normal"`
}

// MarshalJSON json.Marshaller implementation.
func (d DetectorPlane) Validate() error {
	return nil
}
