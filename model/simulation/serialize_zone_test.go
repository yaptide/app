package simulation

import (
	"testing"

	"github.com/yaptide/yaptide/pkg/converter/specs"
	"github.com/yaptide/yaptide/test"
)

var zoneTestCases = []test.SerializeTestCase{
	{
		RawValue: mRaw{
			"id":           int64(0),
			"baseId":       int64(0),
			"parentId":     int64(0),
			"materialId":   int64(0),
			"name":         "",
			"construction": aRaw{},
		},
		ObjectValue: Zone{
			Zone: specs.Zone{
				Construction: []specs.ZoneOperation{},
			},
		},
	},
}

func TestSerializeZones(t *testing.T) {
	test.RunSerializeTestCases(
		t, zoneTestCases,
		zoneMarshaler,
		zoneUnmarshaler,
	)
}
