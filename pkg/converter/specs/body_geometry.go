package specs

import (
	"fmt"

	"github.com/yaptide/yaptide/pkg/converter/geometry"
)

// BodySphere represent sphere with given radius in space.
type BodySphere struct {
	Center geometry.Point
	Radius float64
}

// BodyCuboid represent cuboid of given sizes in a space.
type BodyCuboid struct {
	Center geometry.Point
	Size   geometry.Vec3D
}

// BodyCylinder represent cylinder of given sizes in a space.
type BodyCylinder struct {
	Center geometry.Point
	Height float64
	Radius float64
}

// Validate ...
func (b BodySphere) Validate() error {
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
func (b BodyCuboid) Validate() error {
	result := mErr{}

	if err := b.Size.ValidatePositive(); err != nil {
		result["size"] = err
	}

	if len(result) > 0 {
		return result
	}
	return nil
}

func (b BodyCylinder) Validate() error {
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
