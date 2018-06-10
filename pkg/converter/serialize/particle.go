package serialize

import (
	"fmt"

	"github.com/yaptide/yaptide/pkg/converter/setup"
)

func ParticleMarshaler(p setup.Particle) marshaler {
	return func() (interface{}, error) {
		switch t := p.(type) {
		case setup.PredefinedParticle:
			return mRaw{"type": string(t)}, nil
		case setup.AllParticles:
			return mRaw{"type": "all"}, nil
		case setup.HeavyIon:
			return StructMarshaler(func(m fieldMarshaler) {
				m("type", StringMarshaler("heavy_ion"))
				m("charge", Int64Marshaler(t.Charge))
				m("nucleonsCount", Int64Marshaler(t.NucleonsCount))
			})()
		default:
			log.Errorf("unknown particle type %d", p)
			return nil, fmt.Errorf("unknown particle type")
		}
	}
}

func ParticleUnmarshaler(p *setup.Particle) unmarshaler {
	return UnionTypeUnmarshaler(func(unionType string) unmarshaler {
		switch unionType {
		case "all":
			*p = setup.AllParticles("all")
		case "heavy_ion":
			particle := setup.HeavyIon{}
			*p = particle
			return StructUnmarshaler(func(u fieldUnmarshaler) {
				u("charge", Int64Unmarshaler(&particle.Charge))
				u("nucleonsCount", Int64Unmarshaler(&particle.NucleonsCount))
			})
		default:
			*p = setup.PredefinedParticle(unionType)
		}
		return func(raw interface{}) error {
			return nil
		}
	})
}
