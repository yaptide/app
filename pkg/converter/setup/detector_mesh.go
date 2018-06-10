package setup

import (
	"github.com/yaptide/yaptide/pkg/converter/geometry"
)

// DetectorMesh detector.
type DetectorMesh struct {
	Center geometry.Point    `json:"center"`
	Size   geometry.Vec3D    `json:"size"`
	Slices geometry.Vec3DInt `json:"slices"`
}

// Validate ...
func (d DetectorMesh) Validate() error {
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
