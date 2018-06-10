package setup

// SimulationOptions ...
type SimulationOptions struct {
	// AntyparticleCorrectionOn ...
	AntyparticleCorrectionOn bool `json:"antyparticleCorrectionOn"`
	// NuclearReactionsOn ...
	NuclearReactionsOn bool `json:"nuclearReactionsOn"`
	// MeanEnergyLoss ...
	MeanEnergyLoss Fraction `json:"meanEnergyLoss"`
	// MinEnergyLoss ...
	MinEnergyLoss float64 `json:"minEnergyLoss"`
	// ScatteringType ...
	ScatteringType ScatteringType `json:"scatteringType"`
	// EnergyStraggling ...
	EnergyStraggling EnergyStragglingType `json:"energyStraggling"`
	// FastNeutronTransportOn ...
	FastNeutronTransportOn bool `json:"fastNeutronTransportOn"`
	// LowEnergyNeutronCutOff ...
	LowEnergyNeutronCutOff float64 `json:"lowEnergyNeutronCutOff"`
	// NumberOfGeneratedParticles ...
	NumberOfGeneratedParticles int64 `json:"numberOfGeneratedParticles"`
}

// DefaultOptions ...
var DefaultOptions = SimulationOptions{}

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
