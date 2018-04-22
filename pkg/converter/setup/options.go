package setup

import (
	"encoding/json"
	"fmt"
)

// SimulationOptions ...
type SimulationOptions struct {
	// AntyparticleCorrectionOn ...
	// SHIELD doc: APCORR
	AntyparticleCorrectionOn bool `json:"antyparticleCorrectionOn"`
	// NuclearReactionsOn ...
	// SHIELD doc: NUCRE
	NuclearReactionsOn bool `json:"nuclearReactionsOn"`

	// MeanEnergyLoss ...
	// SHIELD doc: DELTAE
	MeanEnergyLoss Fraction `json:"meanEnergyLoss"`
	// MinEnergyLoss ...
	// SHIELD doc: DEMIN
	MinEnergyLoss float64 `json:"minEnergyLoss"`

	// ScatteringType ...
	// SHIELD doc: MSCAT
	ScatteringType ScatteringType `json:"scatteringType"`
	// EnergyStraggling ...
	// SHIELD doc: STRAGG
	EnergyStraggling EnergyStragglingType `json:"energyStraggling"`

	// FastNeutronTransportOn ...
	// SHIELD doc: NEUTRFAST
	FastNeutronTransportOn bool `json:"fastNeutronTransportOn"`
	// LowEnergyNeutronCutOff ...
	// SHIELD doc: NEUTRLCUT
	LowEnergyNeutronCutOff float64 `json:"lowEnergyNeutronCutOff"`

	// NumberOfGeneratedParticles ...
	// SHIELD doc: NTSTAT
	NumberOfGeneratedParticles int64 `json:"numberOfGeneratedParticles"`
}

// DefaultOptions ...
var DefaultOptions = SimulationOptions{
	AntyparticleCorrectionOn:   false,
	NuclearReactionsOn:         true,
	MeanEnergyLoss:             1,
	MinEnergyLoss:              0.025,
	ScatteringType:             MoliereScattering,
	EnergyStraggling:           VavilovStraggling,
	FastNeutronTransportOn:     true,
	LowEnergyNeutronCutOff:     0,
	NumberOfGeneratedParticles: 1000,
}

// ScatteringType ...
type ScatteringType int64

const (
	// NoScattering ...
	NoScattering ScatteringType = iota
	// GaussianScattering ...
	GaussianScattering
	// MoliereScattering ...
	MoliereScattering
)

var mapScatteringToJSON = map[ScatteringType]string{
	NoScattering:       "",
	GaussianScattering: "gaussian",
	MoliereScattering:  "moliere",
}

var mapJSONToScattering = map[string]ScatteringType{
	"":         NoScattering,
	"gaussian": GaussianScattering,
	"moliere":  MoliereScattering,
}

// MarshalJSON json.Marshaller implementation.
func (s ScatteringType) MarshalJSON() ([]byte, error) {
	res, ok := mapScatteringToJSON[s]
	if !ok {
		return nil,
			fmt.Errorf("ScatteringType.MarshalJSON: can not convert %v to string", s)
	}
	return json.Marshal(res)
}

// UnmarshalJSON json.Unmarshaller implementation.
func (s *ScatteringType) UnmarshalJSON(b []byte) error {
	var input string
	err := json.Unmarshal(b, &input)
	if err != nil {
		return err
	}

	newType, ok := mapJSONToScattering[input]
	if !ok {
		return fmt.Errorf(
			"StateOfMatter.UnmarshalJSON: can not convert %s to StateOfMatter", input)
	}
	*s = newType
	return nil
}

// EnergyStragglingType ...
type EnergyStragglingType int64

const (
	// NoEnergyStraggling ...
	NoEnergyStraggling EnergyStragglingType = iota
	// VavilovStraggling ...
	VavilovStraggling
	// GaussianStraggling ...
	GaussianStraggling
)

var mapEnergyStragglingToJSON = map[EnergyStragglingType]string{
	NoEnergyStraggling: "",
	VavilovStraggling:  "vavilov",
	GaussianStraggling: "gaussian",
}

var mapJSONToEnergyStraggling = map[string]EnergyStragglingType{
	"":         NoEnergyStraggling,
	"vavilov":  VavilovStraggling,
	"gaussian": GaussianStraggling,
}

// MarshalJSON json.Marshaller implementation.
func (s EnergyStragglingType) MarshalJSON() ([]byte, error) {
	res, ok := mapEnergyStragglingToJSON[s]
	if !ok {
		return nil,
			fmt.Errorf("EnergyStragglingType.MarshalJSON: can not convert %v to string", s)
	}
	return json.Marshal(res)
}

// UnmarshalJSON json.Unmarshaller implementation.
func (s *EnergyStragglingType) UnmarshalJSON(b []byte) error {
	var input string
	err := json.Unmarshal(b, &input)
	if err != nil {
		return err
	}

	newType, ok := mapJSONToEnergyStraggling[input]
	if !ok {
		return fmt.Errorf(
			"EnergyStragglingType.UnmarshalJSON: can not convert %s to StateOfMatter", input)
	}
	*s = newType
	return nil
}
