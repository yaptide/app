package setup

import (
	"fmt"
)

// MaterialID ...
type MaterialID int64

// Material defines the zone material that is used in the simulation.
type Material struct {
	ID    MaterialID
	Specs MaterialSpecs
}

// MaterialSpecs ...
type MaterialSpecs interface {
	isMaterialType() bool
}

// Validate ...
func (m MaterialID) Validate() error {
	if m < 0 {
		return fmt.Errorf("Material id needs to be positive integer")
	}
	return nil
}

func (m MaterialPredefined) isMaterialType() bool {
	return true
}

func (m MaterialCompound) isMaterialType() bool {
	return true
}

func (m MaterialVoxel) isMaterialType() bool {
	return true
}
