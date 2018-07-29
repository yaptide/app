// Package geometry ...
package geometry

import (
	"fmt"
	"github.com/yaptide/yaptide/pkg/converter/errors"
)

type mErr = errors.MErr

// Point represent a point in space.
type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

// Vec3DCylindricalInt 3-dimensional vector of integers in cylindrical cordinates.
type Vec3DCylindricalInt struct {
	Radius int64 `json:"radius"`
	Angle  int64 `json:"angle"`
	Z      int64 `json:"z"`
}

// Vec3D represent 3-dimensional vector.
type Vec3D struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

// ValidatePositive ...
func (v Vec3D) ValidatePositive() error {
	result := mErr{}
	fieldErr := fmt.Errorf("should be larger than 0")

	if v.X <= 0 {
		result["x"] = fieldErr
	}

	if v.Y <= 0 {
		result["y"] = fieldErr
	}

	if v.Z <= 0 {
		result["z"] = fieldErr
	}

	if len(result) > 0 {
		return result
	}
	return nil
}

// Vec3DInt represent 3-dimensional vector of integers.
type Vec3DInt struct {
	X int64 `json:"x"`
	Y int64 `json:"y"`
	Z int64 `json:"z"`
}

// ValidatePositive ...
func (v Vec3DInt) ValidatePositive() error {
	result := mErr{}
	fieldErr := fmt.Errorf("should be larger than 0")

	if v.X <= 0 {
		result["x"] = fieldErr
	}

	if v.Y <= 0 {
		result["y"] = fieldErr
	}

	if v.Z <= 0 {
		result["z"] = fieldErr
	}

	if len(result) > 0 {
		return result
	}
	return nil
}

// Range contain min and max value of certain quantity.
type Range struct {
	Min float64 `json:"min"`
	Max float64 `json:"max"`
}

// Validate ...
func (r Range) Validate() error {
	if r.Min > r.Max {
		return mErr{
			"min": fmt.Errorf("can't be larger than maximum value"),
			"max": fmt.Errorf("can't be smaller than minimal value"),
		}
	}
	return nil
}

// ValidatePositive ...
func (r Range) ValidatePositive() error {
	result := r.Validate().(mErr)
	if result == nil {
		result = mErr{}
	}

	if r.Min < 0 {
		result["min"] = fmt.Errorf("can't be negative number")
	}

	if r.Max < 0 {
		result["max"] = fmt.Errorf("can't be negative number")
	}
	if len(result) > 0 {
		return result
	}
	return nil
}

// ValidateFunc ...
func (r Range) ValidateFunc(validate func(float64) error) error {
	result := r.Validate().(mErr)
	if result == nil {
		result = mErr{}
	}

	if err := validate(r.Min); err != nil {
		result["min"] = err
	}

	if err := validate(r.Max); err != nil {
		result["max"] = err
	}

	if len(result) > 0 {
		return result
	}
	return nil
}
