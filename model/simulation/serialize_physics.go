package simulation

import (
	"fmt"

	"github.com/yaptide/yaptide/pkg/converter/specs"
	. "github.com/yaptide/yaptide/pkg/serialize"
)

var mapStateOfMatterToJSON = map[specs.StateOfMatter]string{
	specs.UndefinedStateOfMatter: "",
	specs.Gas:                    "gas",
	specs.Liquid:                 "liquid",
	specs.Solid:                  "solid",
}

var mapJSONToStateOfMatter = map[string]specs.StateOfMatter{
	"":       specs.UndefinedStateOfMatter,
	"gas":    specs.Gas,
	"liquid": specs.Liquid,
	"solid":  specs.Solid,
}

func stateOfMatterMarshaler(o specs.StateOfMatter) marshaler {
	return EnumMarshaler(o, mapStateOfMatterToJSON)
}

func stateOfMatterUnmarshaler(o *specs.StateOfMatter) unmarshaler {
	return EnumUnmarshaler(o, mapJSONToStateOfMatter)
}

func particleMarshaler(p specs.Particle) marshaler {
	return func() (interface{}, error) {
		switch t := p.(type) {
		case specs.PredefinedParticle:
			return mRaw{"type": string(t)}, nil
		case specs.AllParticles:
			return mRaw{"type": "all"}, nil
		case specs.HeavyIon:
			return StructMarshaler(func(m fieldMarshaler) {
				m("type", StringMarshaler("heavy_ion"))
				m("charge", Int64Marshaler(t.Charge))
				m("nucleonsCount", Int64Marshaler(t.NucleonsCount))
			})()
		case nil:
			return nil, nil
		default:
			log.Errorf("unknown particle type %d", p)
			return nil, fmt.Errorf("unknown particle type")
		}
	}
}

func particleUnmarshaler(p *specs.Particle) unmarshaler {
	return UnionTypeUnmarshaler(func(unionType string) unmarshaler {
		switch unionType {
		case "all":
			*p = specs.AllParticles("all")
		case "heavy_ion":
			return StructUnmarshaler(func(u fieldUnmarshaler) {
				particle := specs.HeavyIon{}
				u("charge", Int64Unmarshaler(&particle.Charge))
				u("nucleonsCount", Int64Unmarshaler(&particle.NucleonsCount))
				*p = particle
			})
		default:
			*p = specs.PredefinedParticle(unionType)
		}
		return func(raw interface{}) error {
			return nil
		}
	})
}
