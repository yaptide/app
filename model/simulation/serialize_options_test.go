package simulation

import (
	"testing"

	"github.com/yaptide/yaptide/test"
)

var optionsTestCases = []test.SerializeTestCase{
	{
		RawValue: mRaw{
			"antyparticleCorrectionOn":   false,
			"nuclearReactionsOn":         false,
			"meanEnergyLoss":             float64(0),
			"minEnergyLoss":              float64(0),
			"scatteringType":             "",
			"energyStragglingType":       "",
			"fastNeutronTransportOn":     false,
			"lowEnergyNeutronCutOff":     float64(0),
			"numberOfGeneratedParticles": int64(0),
		},
		ObjectValue: SimulationOptions{},
	},
}

func TestSerializeOptions(t *testing.T) {
	test.RunSerializeTestCases(
		t, optionsTestCases, optionsMarshaler, optionsUnmarshaler,
	)
}
