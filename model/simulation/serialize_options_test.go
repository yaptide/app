package simulation

import (
	"testing"

	"github.com/yaptide/yaptide/pkg/converter/specs"
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
	}, {
		RawValue: mRaw{
			"antyparticleCorrectionOn":   true,
			"nuclearReactionsOn":         true,
			"meanEnergyLoss":             float64(0.1),
			"minEnergyLoss":              float64(0.2),
			"scatteringType":             "gaussian",
			"energyStragglingType":       "vavilov",
			"fastNeutronTransportOn":     true,
			"lowEnergyNeutronCutOff":     float64(0.3),
			"numberOfGeneratedParticles": int64(4),
		},
		ObjectValue: SimulationOptions{
			AntyparticleCorrectionOn:   true,
			NuclearReactionsOn:         true,
			MeanEnergyLoss:             0.1,
			MinEnergyLoss:              0.2,
			ScatteringType:             specs.GaussianScattering,
			EnergyStraggling:           specs.VavilovStraggling,
			FastNeutronTransportOn:     true,
			LowEnergyNeutronCutOff:     0.3,
			NumberOfGeneratedParticles: 4,
		},
	},
}

func TestSerializeOptions(t *testing.T) {
	test.RunSerializeTestCases(
		t, optionsTestCases, optionsMarshaler, optionsUnmarshaler,
	)
}
