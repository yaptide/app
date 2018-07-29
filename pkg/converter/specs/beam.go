package specs

import (
	"github.com/yaptide/yaptide/pkg/converter/geometry"
)

// Beam ...
type Beam struct {
	// Direction ...
	Direction BeamDirection
	// Divergance ...
	Divergence BeamDivergence
	// Particle ...
	Particle Particle
	// InitialBaseEnergy ...
	InitialBaseEnergy float64
	// InitialEnergySigma ...
	InitialEnergySigma float64
}

// DefaultBeam represents default beam configuration.
var DefaultBeam = Beam{}

// BeamDirection ...
type BeamDirection struct {
	// Phi is angle between positive x axis and direction after cast on xy plane.
	Phi float64
	// Theta is angle between z axis and direction.
	Theta    float64
	Position geometry.Point
}

// BeamDivergence ...
type BeamDivergence struct {
	SigmaX       float64      `json:"sigmaX"`
	SigmaY       float64      `json:"sigmaY"`
	Distribution Distribution `json:"distribution"`
}
