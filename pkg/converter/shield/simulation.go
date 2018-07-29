package shield

import (
	"github.com/yaptide/yaptide/pkg/converter/shield/geometry"
	"github.com/yaptide/yaptide/pkg/converter/shield/material"
	"github.com/yaptide/yaptide/pkg/converter/specs"
)

// SerializationContext is struct used to recover data lost in process of
// serializing simulation data.
type SerializationContext struct {
	MapMaterialID           map[material.ShieldID]specs.MaterialID
	MapBodyID               map[geometry.ShieldBodyID]specs.BodyID
	MapFilenameToDetectorID map[string]specs.DetectorID
}

// NewSerializationContext constructor.
func NewSerializationContext() SerializationContext {
	return SerializationContext{
		MapMaterialID:           map[material.ShieldID]specs.MaterialID{},
		MapBodyID:               map[geometry.ShieldBodyID]specs.BodyID{},
		MapFilenameToDetectorID: map[string]specs.DetectorID{},
	}
}
