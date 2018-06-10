package setup

import (
	"fmt"
)

// MaterialCompound material type - create material by defining isotope mixture.
type MaterialCompound struct {
	Name string `json:"name"`

	// Density of the medium in g/cm³ - mandatory.
	Density float64 `json:"density"`

	// StateOfMatter - mandatory.
	StateOfMatter StateOfMatter `json:"stateOfMatter"`

	Elements []Element `json:"elements"`

	// Load stopping power from external file - optional. The file is selected
	// just like in case of Predefined material named by this string.
	ExternalStoppingPowerFromPredefined string `json:"externalStoppingPowerFromPredefined,omitempty"`
}

// Element is a basic building block of Compound.
type Element struct {
	Isotope string `json:"isotope"`

	RelativeStoichiometricFraction int64 `json:"relativeStoichiometricFraction"`

	// Override atomic mass - optional.
	AtomicMass *int64 `json:"atomicMass,omitempty"`

	// Mean excitation energy (I-value) in eV - optional.
	IValue *float64 `json:"iValue,omitempty"`
}

// Validate ...
func (m MaterialCompound) Validate() error {
	result := mErr{}

	if m.Name == "" {
		result["name"] = fmt.Errorf("detector name can't be empty")
	}

	if m.Density <= 0 {
		result["density"] = fmt.Errorf("density needs to be positive number")
	}

	if m.StateOfMatter == UndefinedStateOfMatter {
		result["stateOfMatter"] = fmt.Errorf("state of matter is required")
	}

	if len(m.Elements) == 0 {
		result["elements"] = fmt.Errorf(
			"compound material need to have defined at least one element",
		)
	}
	elementsResult := make(aErr, len(m.Elements))
	elementsHasError := false
	for i, element := range m.Elements {
		err := element.Validate()
		if err != nil {
			elementsHasError = true
		}
		elementsResult[i] = err
	}
	if elementsHasError {
		result["elements"] = elementsResult
	}

	if len(result) > 0 {
		return result
	}
	return nil
}

// Validate ...
func (e Element) Validate() error {
	result := mErr{}

	if _, exists := IsotopesSet[e.Isotope]; !exists {
		result["isotope"] = fmt.Errorf("Unknown isotope %s", e.Isotope)
	}

	if e.RelativeStoichiometricFraction <= 0 {
		result["relativeStoichiometricFraction"] = fmt.Errorf(
			"should be positive integer",
		)
	}

	if e.AtomicMass != nil && *e.AtomicMass <= 0 {
		result["atomicMass"] = fmt.Errorf("should be positive integer")
	}

	if e.IValue != nil && *e.IValue <= 0 {
		result["iValue"] = fmt.Errorf("should be positive")
	}

	if len(result) > 0 {
		return result
	}
	return nil
}
