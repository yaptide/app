package serialize

import (
	"github.com/yaptide/yaptide/pkg/converter/setup"
)

var mapDistributionToJSON = map[setup.Distribution]string{
	setup.NoDistribution:       "",
	setup.FlatDistribution:     "flat",
	setup.GaussianDistribution: "gaussian",
}

var mapJSONToDistribution = func() map[string]setup.Distribution {
	mapping := map[string]setup.Distribution{}
	for key, value := range mapDistributionToJSON {
		mapping[value] = key
	}
	return mapping
}()

func DistributionMarshaler(d setup.Distribution) marshaler {
	return EnumMarshaler(d, mapDistributionToJSON)
}

func DistributionUnmarshaler(d *setup.Distribution) unmarshaler {
	return EnumUnmarshaler(d, mapJSONToDistribution)
}
