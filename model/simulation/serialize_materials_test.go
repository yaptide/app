package simulation

import (
	"testing"

	"github.com/yaptide/yaptide/pkg/converter/specs"
	"github.com/yaptide/yaptide/test"
)

var materialsTestCases = []test.SerializeTestCase{
	{
		RawValue: mRaw{
			"name": "material name",
			"id":   int64(1),
			"specs": mRaw{
				"type":          "predefined",
				"predefinedId":  "",
				"density":       nil,
				"stateOfMatter": "",
			},
		},
		ObjectValue: Material{
			Name: "material name",
			Material: specs.Material{
				ID:    specs.MaterialID(1),
				Specs: specs.MaterialPredefined{},
			},
		},
	},
	{
		RawValue: mRaw{
			"name": "material name",
			"id":   int64(1),
			"specs": mRaw{
				"type":          "compound",
				"density":       float64(0),
				"stateOfMatter": "",
				"elements":      []interface{}{},
			},
		},
		ObjectValue: Material{
			Name: "material name",
			Material: specs.Material{
				ID:    specs.MaterialID(1),
				Specs: specs.MaterialCompound{},
			},
		},
	},
}

func TestSerializeMaterials(t *testing.T) {
	test.RunSerializeTestCases(
		t, materialsTestCases,
		materialMarshaler,
		materialUnmarshaler,
	)
}
