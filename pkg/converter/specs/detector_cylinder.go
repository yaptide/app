package specs

import (
	"github.com/yaptide/yaptide/pkg/converter/geometry"
	"github.com/yaptide/yaptide/pkg/converter/validate"
)

// DetectorCylinder is detector with cylindrical shape directed along z-axis.
type DetectorCylinder struct {
	Radius geometry.Range
	Angle  geometry.Range
	ZValue geometry.Range
	Slices geometry.Vec3DCylindricalInt
}

// Validate ...
func (d DetectorCylinder) Validate() error {
	result := mErr{}

	if err := d.Radius.ValidatePositive(); err != nil {
		result["radius"] = err
	}

	if err := d.Angle.ValidateFunc(validate.InRange2PI); err != nil {
		result["angle"] = err
	}

	if err := d.ZValue.Validate(); err != nil {
		result["zValue"] = err
	}

	if len(result) > 0 {
		return result
	}
	return nil
}