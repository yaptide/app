package setup

import (
	"fmt"
)

var predefinedParticleTypes = map[string]bool{
	"neutron":          true,
	"proton":           true,
	"pion_pi_minus":    true,
	"pion_pi_plus":     true,
	"pion_pi_zero":     true,
	"he_3":             true,
	"he_4":             true,
	"anti_neutron":     true,
	"anti_proton":      true,
	"kaon_minus":       true,
	"kaon_plus":        true,
	"kaon_zero":        true,
	"kaon_anti":        true,
	"gamma":            true,
	"electron":         true,
	"positron":         true,
	"muon_minus":       true,
	"muon_plus":        true,
	"e_neutrino":       true,
	"e_anti_neutrino":  true,
	"mi_neutrino":      true,
	"mi_anti_neutrino": true,
	"deuteron":         true,
	"triton":           true,
}

// Particle is interface for particle scored in detectors.
type Particle interface {
	isParticle() bool
}

func (p AllParticles) isParticle() bool {
	return true
}

func (p PredefinedParticle) isParticle() bool {
	return true
}

func (p HeavyIon) isParticle() bool {
	return true
}

// AllParticles ...
type AllParticles string

// PredefinedParticle ...
type PredefinedParticle string

// HeavyIon ...
type HeavyIon struct {
	Charge        int64 `json:"charge"`
	NucleonsCount int64 `json:"nucleonsCount"`
}

// Validate ...
func (p PredefinedParticle) Validate() error {
	_, exists := predefinedParticleTypes[string(p)]
	if !exists {
		return fmt.Errorf("%v is not a predefined particle type", p)
	}
	return nil
}

// Validate ...
func (p HeavyIon) Validate() error {
	result := mErr{}
	if p.Charge <= 2 {
		result["charge"] = fmt.Errorf("Number of protons must be larger than 2")
	}
	if p.Charge > p.NucleonsCount && p.NucleonsCount > 0 {
		result["charge"] = fmt.Errorf("Number of protons can't be larger than number of nucleons")
	}
	if p.NucleonsCount <= 0 {
		result["nucleonsCount"] = fmt.Errorf("Number of nucleons must be larger than 0")
	}
	return result
}
