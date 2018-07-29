package shield

import (
	"github.com/yaptide/yaptide/pkg/converter/shield/detector"
	"github.com/yaptide/yaptide/pkg/converter/shield/geometry"
	"github.com/yaptide/yaptide/pkg/converter/shield/material"
	"github.com/yaptide/yaptide/pkg/converter/specs"
)

// RawShieldSpecs is input for shield Serialize function.
type RawShieldSpecs struct {
	Materials material.Materials
	Geometry  geometry.Geometry
	Detectors []detector.Detector
	Beam      specs.Beam
	Options   specs.SimulationOptions
}

// Files ...
func (s RawShieldSpecs) Files() map[string]string {
	return SerializeData(s)
}
