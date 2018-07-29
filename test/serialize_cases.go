package test

import (
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/require"
)

type SerializeTestCase struct {
	RawValue    interface{}
	ObjectValue interface{}
}

// RunSerializeTestCases
// !!! there is no type checking, arguments of this function are not validated use with caution. !!!
func RunSerializeTestCases(
	t *testing.T, testCases []SerializeTestCase, marshaler interface{}, unmarshaler interface{},
) {
	m := reflect.ValueOf(marshaler)
	u := reflect.ValueOf(unmarshaler)

	for _, testCase := range testCases {
		runSerializeTestCase(t, testCase, m, u)
	}
}

func runSerializeTestCase(
	t *testing.T, testCase SerializeTestCase, marshaler reflect.Value, unmarshaler reflect.Value,
) {
	runSerializeTestCaseMarshal(t, testCase, marshaler)
	runSerializeTestCaseUnmarshal(t, testCase, unmarshaler)
	runSerializeTestCaseMarshalUnmarshal(t, testCase, marshaler, unmarshaler)
	runSerializeTestCaseUnmarshalMarshal(t, testCase, marshaler, unmarshaler)
}

func runSerializeTestCaseMarshal(
	t *testing.T, testCase SerializeTestCase, marshaler reflect.Value,
) {
	t.Logf("runSerializeTestCaseMarshal %T", testCase.ObjectValue)
	returnList := marshaler.Call([]reflect.Value{
		reflect.ValueOf(testCase.ObjectValue),
	})[0].Call([]reflect.Value{})
	require.Len(t, returnList, 2)
	require.Nil(
		t, returnList[1].Interface(),
		spew.Sdump(returnList[1].Interface()),
	)
	AssertDeepEqual(t, testCase.RawValue, returnList[0].Interface())
}

func runSerializeTestCaseUnmarshal(
	t *testing.T, testCase SerializeTestCase, unmarshaler reflect.Value,
) {
	t.Logf("runSerializeTestCaseUnmarshal %T", testCase.ObjectValue)
	unpack := reflect.New(unmarshaler.Type().In(0).Elem())
	returnList := unmarshaler.Call([]reflect.Value{
		unpack,
	})[0].Call([]reflect.Value{
		reflect.ValueOf(testCase.RawValue),
	})
	require.Len(t, returnList, 1)
	require.Nil(
		t, returnList[0].Interface(),
		spew.Sdump(returnList[0].Interface()),
	)
	AssertDeepEqual(t, testCase.ObjectValue, unpack.Elem().Interface())
}

func runSerializeTestCaseMarshalUnmarshal(
	t *testing.T, testCase SerializeTestCase, marshaler reflect.Value, unmarshaler reflect.Value,
) {
}

func runSerializeTestCaseUnmarshalMarshal(
	t *testing.T, testCase SerializeTestCase, marshaler reflect.Value, unmarshaler reflect.Value,
) {
}
