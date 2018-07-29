package specs

import (
	"fmt"

	"github.com/yaptide/yaptide/pkg/converter/validate"
)

// Validate ...
func (b Beam) Validate() error {
	result := mErr{}

	if b.InitialBaseEnergy < 0 {
		result["initialBaseEnergy"] = fmt.Errorf("shuld be positive value")
	}

	if b.InitialEnergySigma < 0 {
		result["initialEnergySigma"] = fmt.Errorf("should be positive value")
	}

	return result
}

// Validate ...
func (b BeamDirection) Validate() error {
	result := mErr{}

	if err := validate.InRange2PI(b.Phi); err != nil {
		result["phi"] = err
	}
	if err := validate.InRangePI(b.Theta); err != nil {
		result["theta"] = err
	}

	if len(result) > 0 {
		return result
	}
	return nil
}

// Validate ...
func (b BeamDivergence) Validate() error {
	result := mErr{}

	// TODO research this better;
	return result
}
