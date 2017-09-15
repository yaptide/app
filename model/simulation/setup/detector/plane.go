package detector

import (
	"encoding/json"
	"github.com/Palantir/palantir/model/simulation/common"
)

// Plane detector.
type Plane struct {
	Point  common.Point
	Normal common.Vec3D
	Slices common.Vec3DInt
}

// MarshalJSON json.Marshaller implementation.
func (g Plane) MarshalJSON() ([]byte, error) {
	type Alias Plane
	return json.Marshal(struct {
		detectorType
		Alias
	}{
		detectorType: planeScoringDetector,
		Alias:        (Alias)(g),
	})
}