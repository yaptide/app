package setup

import (
	"encoding/json"
	"testing"

	test "github.com/yaptide/converter/test"
)

var opTestCases = test.MarshallingCases{
	{
		&Operation{BodyID: ID(1), Type: Intersect},
		`{"bodyId":1,"type":"intersect"}`,
	},
	{
		&Operation{BodyID: ID(1), Type: Subtract},
		`{"bodyId":1,"type":"subtract"}`,
	},
	{
		&Operation{BodyID: ID(1), Type: Union},
		`{"bodyId":1,"type":"union"}`,
	},
}

func TestOperationMarshal(t *testing.T) {
	test.Marshal(t, opTestCases)
}

func TestOperationUnmarshal(t *testing.T) {
	test.Unmarshal(t, opTestCases)
}

func TestOperationMarshalUnmarshalled(t *testing.T) {
	test.MarshalUnmarshalled(t, opTestCases)
}

func TestOperationUnmarshalMarshalled(t *testing.T) {
	test.UnmarshalMarshalled(t, opTestCases)
}

func TestOperationInvalidTypeMarshal(t *testing.T) {
	testCases := []struct {
		TestOperation *Operation
		IsReturnErr   bool
	}{
		{
			&Operation{BodyID: ID(1), Type: Subtract},
			false,
		},
		{
			&Operation{BodyID: ID(1), Type: (OperationType)(10000)},
			true,
		},
	}

	for _, tc := range testCases {
		_, err := json.Marshal(tc.TestOperation)
		if (err != nil) != tc.IsReturnErr {
			t.Errorf("TestOperationInvalidTypeMarshal: IsReturnErr: %v, Actual: %v",
				tc.IsReturnErr, !tc.IsReturnErr)
		}
	}
}

func TestOperationInvalidTypeUnmarshal(t *testing.T) {
	testCases := []struct {
		TestJSON    string
		IsReturnErr bool
	}{
		{
			`{"bodyId":1,"type":"intersect"}`,
			false,
		},
		{
			`{"bodyId":1,"type":"xxxxxxx"}`,
			true,
		},
	}

	for _, tc := range testCases {
		var op Operation
		err := json.Unmarshal([]byte(tc.TestJSON), &op)
		if (err != nil) != tc.IsReturnErr {
			t.Errorf("TestOperationInvalidTypeUnmarshal: IsReturnErr: %v, Actual: %v",
				tc.IsReturnErr, !tc.IsReturnErr)
		}
	}
}
