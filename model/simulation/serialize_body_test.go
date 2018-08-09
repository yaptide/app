package simulation

import (
	"testing"

	"github.com/yaptide/yaptide/pkg/converter/specs"
	"github.com/yaptide/yaptide/test"
)

var bodyTestCases = []test.SerializeTestCase{
	{
		RawValue: mRaw{
			"id": int64(0),
			"geometry": mRaw{
				"type": "cuboid",
				"center": mRaw{
					"x": float64(0),
					"y": float64(0),
					"z": float64(0),
				},
				"size": mRaw{
					"x": float64(0),
					"y": float64(0),
					"z": float64(0),
				},
			},
		},
		ObjectValue: specs.Body{
			Geometry: specs.BodyCuboid{},
		},
	},
	{
		RawValue: mRaw{
			"id": int64(0),
			"geometry": mRaw{
				"type": "sphere",
				"center": mRaw{
					"x": float64(0),
					"y": float64(0),
					"z": float64(0),
				},
				"radius": float64(0),
			},
		},
		ObjectValue: specs.Body{
			Geometry: specs.BodySphere{},
		},
	},
	{
		RawValue: mRaw{
			"id": int64(0),
			"geometry": mRaw{
				"type": "cylinder",
				"center": mRaw{
					"x": float64(0),
					"y": float64(0),
					"z": float64(0),
				},
				"radius": float64(0),
				"height": float64(0),
			},
		},
		ObjectValue: specs.Body{
			Geometry: specs.BodyCylinder{},
		},
	},
}

func TestSerializeBodies(t *testing.T) {
	test.RunSerializeTestCases(
		t, bodyTestCases,
		bodyMarshaler,
		bodyUnmarshaler,
	)
}
