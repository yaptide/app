package validate

import (
	"fmt"
	"math"
)

const floatingPointTolerance = 0.000001

// IsInRange ...
func IsInRange(start float64, end float64, value float64) bool {
	return value >= start && value <= end
}

// IsInRange2PI ...
func IsInRange2PI(value float64) bool {
	return value >= 0 && value <= math.Pi*2+floatingPointTolerance
}

// InRange2PI ...
func InRange2PI(value float64) error {
	if IsInRange2PI(value) {
		return fmt.Errorf("should be between 0 and 2PI")
	}
	return nil
}

// IsInRangePI ...
func IsInRangePI(value float64) bool {
	return value >= 0 && value <= math.Pi+floatingPointTolerance
}

// InRangePI ...
func InRangePI(value float64) error {
	if IsInRangePI(value) {
		return fmt.Errorf("should be between 0 and PI")
	}
	return nil
}
