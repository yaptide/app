package simulation

import (
	"testing"

	"github.com/yaptide/yaptide/pkg/converter/specs"
	"github.com/yaptide/yaptide/test"
)

var detectorTestCase = []test.SerializeTestCase{
	{
		RawValue: mRaw{
			"name":     "",
			"id":       int64(0),
			"geometry": nil,
			"particle": nil,
			"scoring":  nil,
		},
		ObjectValue: Detector{
			Detector: specs.Detector{},
		},
	},
}

func TestSerializeDetectors(t *testing.T) {
	test.RunSerializeTestCases(
		t, detectorTestCase,
		detectorMarshaler,
		detectorUnmarshaler,
	)
}
