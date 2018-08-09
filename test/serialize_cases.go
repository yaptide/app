package test

import (
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/require"
	diff "github.com/yudai/gojsondiff"
	"github.com/yudai/gojsondiff/formatter"
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

func RunUnmarshalTestCases(
	t *testing.T, testCases []SerializeTestCase, unmarshaler interface{},
) {
	u := reflect.ValueOf(unmarshaler)

	for _, testCase := range testCases {
		runSerializeTestCaseUnmarshal(t, testCase, u)
	}
}

func RunJsonMarshallTestCase(
	t *testing.T, expectedJson []byte, object interface{},
) {
	actualJson, marshalErr := json.MarshalIndent(object, "  ", "  ")
	require.Nil(t, marshalErr)
	assertRawJsonEqual(t, expectedJson, actualJson)
}

func RunJsonUnmarshallTestCase(
	t *testing.T, expectedObject interface{}, rawJson []byte,
) {
	unpackPtrValue := reflect.New(reflect.TypeOf(expectedObject))
	unmarshalErr := json.Unmarshal(rawJson, unpackPtrValue.Interface())
	require.Nil(t, unmarshalErr)
	AssertDeepEqual(t, expectedObject, unpackPtrValue.Elem().Interface())
}

func RunBsonMarshallTestCase(
	t *testing.T, expectedBson interface{}, object interface{},
) {
	actualBsonRaw, marshalErr := bson.Marshal(object)
	require.Nil(t, marshalErr)
	var actualBson map[string]interface{}
	require.Nil(t, bson.Unmarshal(actualBsonRaw, &actualBson))
	AssertDeepEqual(t, expectedBson, actualBson)
}

func RunBsonUnmarshallTestCase(
	t *testing.T, expectedObject interface{}, bsonMap interface{},
) {
	rawBson, rawBsonErr := bson.Marshal(bsonMap)
	require.Nil(t, rawBsonErr)
	unpackPtrValue := reflect.New(reflect.TypeOf(expectedObject))
	unmarshalErr := bson.Unmarshal(rawBson, unpackPtrValue.Interface())
	require.Nil(t, unmarshalErr)
	AssertDeepEqual(t, expectedObject, unpackPtrValue.Elem().Interface())
}

func runSerializeTestCase(
	t *testing.T, testCase SerializeTestCase, marshaler reflect.Value, unmarshaler reflect.Value,
) {
	runSerializeTestCaseMarshal(t, testCase, marshaler)
	runSerializeTestCaseUnmarshal(t, testCase, unmarshaler)
	runSerializeTestCaseMarshalUnmarshal(t, testCase, marshaler, unmarshaler)
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

// DiffJSON ...
func assertRawJsonEqual(t *testing.T, expected, actual []byte) {
	t.Helper()
	var expectedObject map[string]interface{}
	var actualObject map[string]interface{}
	require.Nil(t, json.Unmarshal(expected, &expectedObject))
	require.Nil(t, json.Unmarshal(actual, &actualObject))

	if !reflect.DeepEqual(expectedObject, actualObject) {
		diffs, diffErr := diff.New().Compare(expected, actual)
		require.Nil(t, diffErr)
		require.True(t, diffs.Modified())

		formatedStr, formatterErr := formatter.NewAsciiFormatter(
			expectedObject, jsonFormatterConfig,
		).Format(diffs)
		require.Nil(t, formatterErr)
		t.Logf("expected json != actual json\n%s", formatedStr)
		t.Fail()
	}
}
