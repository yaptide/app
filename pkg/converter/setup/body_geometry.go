package setup

import (
	"fmt"

	"github.com/yaptide/yaptide/pkg/converter/geometry"
)

// SphereBody represent sphere with given radius in space.
type SphereBody struct {
	Center geometry.Point `json:"center"`
	Radius float64        `json:"radius"`
}

// CuboidBody represent cuboid of given sizes in a space.
type CuboidBody struct {
	Center geometry.Point `json:"center"`
	Size   geometry.Vec3D `json:"size"`
}

// CylinderBody represent cylinder of given sizes in a space.
type CylinderBody struct {
	Center geometry.Point `json:"baseCenter"`
	Height float64        `json:"height"`
	Radius float64        `json:"radius"`
}

// Validate ...
func (b SphereBody) Validate() error {
	result := mErr{}

	if b.Radius <= 0 {
		result["radius"] = fmt.Errorf("should be positive non-zero value")
	}

	if len(result) > 0 {
		return result
	}
	return nil
}

// Validate ...
func (b CuboidBody) Validate() error {
	result := mErr{}

	if err := b.Size.ValidatePositive(); err != nil {
		result["size"] = err
	}

	if len(result) > 0 {
		return result
	}
	return nil
}

func (b CylinderBody) Validate() error {
	result := mErr{}

	if b.Height <= 0 {
		result["height"] = fmt.Errorf("should positive non-zero value")
	}
	if b.Height <= 0 {
		result["radius"] = fmt.Errorf("should positive non-zero value")
	}

	if len(result) > 0 {
		return result
	}
	return nil
}
