package serialize

import (
	"github.com/yaptide/yaptide/pkg/converter/setup"
)

var scateringTypeToJSON = map[setup.ScatteringType]string{
	setup.NoScattering:       "",
	setup.GaussianScattering: "gaussian",
	setup.MoliereScattering:  "moliere",
}

var jsonToScattering = map[string]setup.ScatteringType{
	"":         setup.NoScattering,
	"gaussian": setup.GaussianScattering,
	"moliere":  setup.MoliereScattering,
}

var energyStragglingToJSON = map[setup.EnergyStragglingType]string{
	setup.NoEnergyStraggling: "",
	setup.GaussianStraggling: "gaussian",
	setup.VavilovStraggling:  "vavilov",
}

var jsonToStraggling = map[string]setup.EnergyStragglingType{
	"":         setup.NoEnergyStraggling,
	"gaussian": setup.GaussianStraggling,
	"vavilov":  setup.VavilovStraggling,
}

func OptionsMarshaler(o setup.SimulationOptions) marshaler {
	return StructMarshaler(func(m fieldMarshaler) {
		m("antyparticleCorrectionOn", BooleanMarshaler(o.AntyparticleCorrectionOn))
		m("nuclearReactionsOn", BooleanMarshaler(o.NuclearReactionsOn))
		m("meanEnergyLoss", Float64Marshaler(float64(o.MeanEnergyLoss)))
		m("minEnergyLoss", Float64Marshaler(o.MinEnergyLoss))
		m("scatteringType", EnumMarshaler(o.ScatteringType, scateringTypeToJSON))
		m("energyStragglingType",
			EnumMarshaler(o.EnergyStraggling, energyStragglingToJSON),
		)
		m("fastNeutronTransportOn", BooleanMarshaler(o.FastNeutronTransportOn))
		m("numberOfGeneratedParticles", Int64Marshaler(o.NumberOfGeneratedParticles))
	})
}

func OptionsUnmarshaler(o *setup.SimulationOptions) unmarshaler {
	return StructUnmarshaler(func(u fieldUnmarshaler) {
		u("antyparticleCorrectionOn", BooleanUnmarshaler(&o.AntyparticleCorrectionOn))
		u("nuclearReactionsOn", BooleanUnmarshaler(&o.NuclearReactionsOn))
		u("meanEnergyLoss", Float64Unmarshaler((*float64)(&o.MeanEnergyLoss)))
		u("minEnergyLoss", Float64Unmarshaler(&o.MinEnergyLoss))
		u("scatteringType", EnumUnmarshaler(&o.ScatteringType, jsonToScattering))
		u("energyStragglingType", EnumUnmarshaler(&o.EnergyStraggling, jsonToScattering))
		u("fastNeutronTransportOn", BooleanUnmarshaler(&o.FastNeutronTransportOn))
		u("numberOfGeneratedParticles", Int64Unmarshaler(&o.NumberOfGeneratedParticles))
	})
}
