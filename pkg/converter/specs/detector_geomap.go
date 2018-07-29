package specs

import (
	"github.com/yaptide/yaptide/pkg/converter/geometry"
)

// DetectorGeomap detector used to debug geometry.
type DetectorGeomap struct {
	Center geometry.Point
	Size   geometry.Vec3D
	Slices geometry.Vec3DInt
}

// Validate ...
func (d DetectorGeomap) Validate() error {
	result := mErr{}
	if err := d.Size.ValidatePositive(); err != nil {
		result["size"] = err
	}

	if err := d.Slices.ValidatePositive(); err != nil {
		result["slices"] = err
	}

	if len(result) > 0 {
		return result
	}
	return nil
}
