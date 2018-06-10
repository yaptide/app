package mapping

import (
	"testing"

	"github.com/yaptide/yaptide/pkg/converter/setup"
)

func TestPredefinedMaterialsToShieldICRUMapping(t *testing.T) {
	for predefinedMaterial := range setup.PredefinedMaterialsSet {
		_, found := PredefinedMaterialsToShieldICRU[predefinedMaterial]
		if !found {
			t.Errorf(
				"PredefinedMaterial mapping to Shield ICRU for \"%s\" not found",
				predefinedMaterial,
			)
		}
	}
}

func TestIsotopeToShieldNUCLIDMapping(t *testing.T) {
	for isotope := range setup.IsotopesSet {
		_, found := IsotopesToShieldNUCLID[isotope]
		if !found {
			t.Errorf(
				"Isotope mapping to Shield NUCLID for \"%s\" not found",
				isotope,
			)
		}
	}
}
