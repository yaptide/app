package setup

import (
	"fmt"
)

// MaterialPredefined material type - choose material definition
// from predefined material list by name.
type MaterialPredefined struct {
	PredefinedID string `json:"predefinedId"`

	// Density of the medium in g/cm³ - optional.
	Density float64 `json:"density,omitempty"`

	// State of matter - optional
	StateOfMatter StateOfMatter `json:"stateOfMatter,omitempty"`

	// Load stopping power from external file.
	LoadExternalStoppingPower bool `json:"loadExternalStoppingPower,omitempty"`
}

// Validate ...
func (m MaterialPredefined) Validate() error {
	result := mErr{}

	if _, exists := PredefinedMaterialsSet[m.PredefinedID]; !exists {
		result["predefinedId"] = fmt.Errorf("Unknown predefined material")
	}

	if m.Density < 0 {
		result["density"] = fmt.Errorf("density can't be negative number")
	}

	if len(result) > 0 {
		return result
	}
	return nil
}
