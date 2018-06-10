package setup

import (
	"github.com/yaptide/yaptide/pkg/converter/geometry"
)

// Beam ...
type Beam struct {
	// Direction ...
	Direction BeamDirection `json:"direction"`
	// Divergance ...
	Divergence BeamDivergence `json:"divergence"`
	// Particle ...
	Particle Particle `json:"particle"`
	// InitialBaseEnergy ...
	InitialBaseEnergy float64 `json:"initialBaseEnergy"`
	// InitialEnergySigma ...
	InitialEnergySigma float64 `json:"initialEnergySigma"`
}

// DefaultBeam represents default beam configuration.
var DefaultBeam = Beam{}

// BeamDirection ...
type BeamDirection struct {
	// Phi is angle between positive x axis and direction after cast on xy plane.
	Phi float64 `json:"phi"`
	// Theta is angle between z axis and direction.
	Theta    float64        `json:"theta"`
	Position geometry.Point `json:"position"`
}

// BeamDivergence ...
type BeamDivergence struct {
	SigmaX       float64      `json:"sigmaX"`
	SigmaY       float64      `json:"sigmaY"`
	Distribution Distribution `json:"distribution"`
}
