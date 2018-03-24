package shield

import (
	"github.com/yaptide/converter/setup/body"
	"github.com/yaptide/converter/setup/detector"
	"github.com/yaptide/converter/setup/material"
)

// MaterialID used directly in shield input files.
type MaterialID int

// BodyID used directly in shield input files.
type BodyID int

// ZoneID used directly in shield input files.
type ZoneID int

// SerializationContext is struct used to recover data lost in process of serializing simulation data.
type SerializationContext struct {
	MapMaterialID           map[MaterialID]material.ID
	MapBodyID               map[BodyID]body.ID
	MapFilenameToDetectorID map[string]detector.ID
}

// NewSerializationContext constructor.
func NewSerializationContext() *SerializationContext {
	return &SerializationContext{
		MapMaterialID:           map[MaterialID]material.ID{},
		MapBodyID:               map[BodyID]body.ID{},
		MapFilenameToDetectorID: map[string]detector.ID{},
	}
}
