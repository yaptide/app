package setup

import (
	"fmt"
)

// DetectorZones ...
type DetectorZones struct {
	Zones []ZoneID `json:"zones"`
}

// Validate ...
func (d DetectorZones) Validate() error {
	if len(d.Zones) == 0 {
		return fmt.Errorf("list of zones can't be empty")
	}
	return nil
}
