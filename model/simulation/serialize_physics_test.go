package simulation

import (
	"github.com/yaptide/yaptide/pkg/converter/specs"
	"github.com/yaptide/yaptide/test"
	"testing"
)

var particleTestCases = []test.SerializeTestCase{
	{
		RawValue:    mRaw{"type": "neutron"},
		ObjectValue: specs.PredefinedParticle("neutron"),
	},
	{
		RawValue:    mRaw{"type": "all"},
		ObjectValue: specs.AllParticles("all"),
	},
	{
		RawValue: mRaw{
			"type":          "heavy_ion",
			"charge":        int64(10),
			"nucleonsCount": int64(11),
		},
		ObjectValue: specs.HeavyIon{
			Charge:        int64(10),
			NucleonsCount: int64(11),
		},
	},
}

func TestSerializeParticle(t *testing.T) {
	test.RunSerializeTestCases(t, particleTestCases, particleMarshaler, particleUnmarshaler)
}
