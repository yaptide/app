package simulation

import (
	"github.com/yaptide/yaptide/pkg/converter/specs"
	. "github.com/yaptide/yaptide/pkg/serialize"
)

var mapDistributionToJSON = map[specs.Distribution]string{
	specs.NoDistribution:       "",
	specs.FlatDistribution:     "flat",
	specs.GaussianDistribution: "gaussian",
}

var mapJSONToDistribution = func() map[string]specs.Distribution {
	mapping := map[string]specs.Distribution{}
	for key, value := range mapDistributionToJSON {
		mapping[value] = key
	}
	return mapping
}()

func distributionMarshaler(d specs.Distribution) marshaler {
	return EnumMarshaler(d, mapDistributionToJSON)
}

func distributionUnmarshaler(d *specs.Distribution) unmarshaler {
	return EnumUnmarshaler(d, mapJSONToDistribution)
}
