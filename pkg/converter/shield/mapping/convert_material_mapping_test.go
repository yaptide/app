package mapping

import (
	"testing"

	"github.com/yaptide/yaptide/pkg/converter/specs"
)

func TestPredefinedMaterialsToShieldICRUMapping(t *testing.T) {
	for predefinedMaterial := range specs.PredefinedMaterialsSet {
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
	for isotope := range specs.IsotopesSet {
		_, found := IsotopesToShieldNUCLID[isotope]
		if !found {
			t.Errorf(
				"Isotope mapping to Shield NUCLID for \"%s\" not found",
				isotope,
			)
		}
	}
}
