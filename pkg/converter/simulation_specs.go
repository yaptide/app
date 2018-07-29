package converter

import (
	"github.com/yaptide/yaptide/pkg/converter/specs"
)

type Material = specs.Material
type Body = specs.Body
type Zone = specs.Zone
type Detector = specs.Detector
type Beam = specs.Beam
type SimulationOptions = specs.SimulationOptions

// Specs contains all simulation data.
type Specs struct {
	Materials []Material
	Bodies    []Body
	Zones     []Zone
	Detectors []Detector
	Beam      Beam
	Options   SimulationOptions
}

// NewEmptySpecs constructor.
func NewEmptySpecs() Specs {
	return Specs{
		Materials: []Material{},
		Bodies:    []Body{},
		Zones:     []Zone{},
		Detectors: []Detector{},
		Beam:      specs.DefaultBeam,
		Options:   specs.DefaultOptions,
	}
}
