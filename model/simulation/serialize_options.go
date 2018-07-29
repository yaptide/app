package simulation

import (
	"github.com/yaptide/yaptide/pkg/converter/specs"
	. "github.com/yaptide/yaptide/pkg/serialize"
)

var scateringTypeToJSON = map[specs.ScatteringType]string{
	specs.NoScattering:       "",
	specs.GaussianScattering: "gaussian",
	specs.MoliereScattering:  "moliere",
}

var jsonToScattering = map[string]specs.ScatteringType{
	"":         specs.NoScattering,
	"gaussian": specs.GaussianScattering,
	"moliere":  specs.MoliereScattering,
}

var energyStragglingToJSON = map[specs.EnergyStragglingType]string{
	specs.NoEnergyStraggling: "",
	specs.GaussianStraggling: "gaussian",
	specs.VavilovStraggling:  "vavilov",
}

var jsonToStraggling = map[string]specs.EnergyStragglingType{
	"":         specs.NoEnergyStraggling,
	"gaussian": specs.GaussianStraggling,
	"vavilov":  specs.VavilovStraggling,
}

func optionsMarshaler(o SimulationOptions) marshaler {
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
		m("lowEnergyNeutronCutOff", Float64Marshaler(o.LowEnergyNeutronCutOff))
		m("numberOfGeneratedParticles", Int64Marshaler(o.NumberOfGeneratedParticles))
	})
}

func optionsUnmarshaler(o *SimulationOptions) unmarshaler {
	return StructUnmarshaler(func(u fieldUnmarshaler) {
		u("antyparticleCorrectionOn", BooleanUnmarshaler(&o.AntyparticleCorrectionOn))
		u("nuclearReactionsOn", BooleanUnmarshaler(&o.NuclearReactionsOn))
		u("meanEnergyLoss", Float64Unmarshaler((*float64)(&o.MeanEnergyLoss)))
		u("minEnergyLoss", Float64Unmarshaler(&o.MinEnergyLoss))
		u("scatteringType", EnumUnmarshaler(&o.ScatteringType, jsonToScattering))
		u("energyStragglingType", EnumUnmarshaler(&o.EnergyStraggling, jsonToStraggling))
		u("fastNeutronTransportOn", BooleanUnmarshaler(&o.FastNeutronTransportOn))
		u("lowEnergyNeutronCutOff", Float64Unmarshaler(&o.LowEnergyNeutronCutOff))
		u("numberOfGeneratedParticles", Int64Unmarshaler(&o.NumberOfGeneratedParticles))
	})
}
