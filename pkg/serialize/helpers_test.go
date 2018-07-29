package serialize

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type marshalTestCase struct {
	value marshaler
	raw   interface{}
}

func TestUtilStructMarshaler(t *testing.T) {
	rawMap, err := StructMarshaler(func(m fieldMarshaler) {
		m("field1", func() (interface{}, error) {
			return "value1", nil
		})
		m("field2", func() (interface{}, error) {
			return 2, nil
		})
	})()
	assert.Nil(t, err)
	assert.IsType(t, map[string]interface{}{}, rawMap)
}

func TestUtilStructUnmarshaler(t *testing.T) {
	type someStruct struct {
		field1 string
		field2 int
	}
	var unpack someStruct
	err := StructUnmarshaler(func(u fieldUnmarshaler) {
		u("field1", func(raw interface{}) error {
			assert.IsType(t, "", raw)
			unpack.field1 = raw.(string)
			return nil
		})
		u("field2", func(raw interface{}) error {
			assert.IsType(t, int(0), raw)
			unpack.field2 = raw.(int)
			return nil
		})
	})(map[string]interface{}{
		"field1": "value1",
		"field2": int(2),
	})
	assert.Nil(t, err)
	assert.Equal(t, "value1", unpack.field1)
	assert.Equal(t, int(2), unpack.field2)
}

func TestUtilListMarshaler(t *testing.T) {
	type listElement string
	inputList := []listElement{"123", "234", "445"}
	rawList, err := ListMarshaler(inputList, func(item listElement) marshaler {
		return func() (interface{}, error) {
			return string(item), nil
		}
	})()

	assert.Nil(t, err)
	assert.IsType(t, []interface{}{}, rawList)
	require.Len(t, rawList, 3)
	assert.Equal(t, "123", rawList.([]interface{})[0])
	assert.Equal(t, "234", rawList.([]interface{})[1])
	assert.Equal(t, "445", rawList.([]interface{})[2])

}

func TestUtilListUnmarshaler(t *testing.T) {
	type listElement string
	inputList := []interface{}{"123", "234", "445"}
	var unmarshaledList []listElement
	err := ListUnmarshaler(&unmarshaledList, func(item *listElement) unmarshaler {
		return func(raw interface{}) error {
			assert.IsType(t, "", raw)
			*item = listElement(raw.(string))
			return nil
		}
	})(inputList)

	assert.Nil(t, err)
	assert.Equal(t, listElement("123"), unmarshaledList[0])
	assert.Equal(t, listElement("234"), unmarshaledList[1])
	assert.Equal(t, listElement("445"), unmarshaledList[2])
}

func TestUtilFloat64Marshaler(t *testing.T) {
	var input float64 = 12.12

	result, err := Float64Marshaler(input)()

	assert.Nil(t, err)
	assert.Equal(t, float64(12.12), result)
}

func TestUtilFloat64Unmarshaler(t *testing.T) {
	var rawInput float64 = 12.12
	var result float64
	err := Float64Unmarshaler(&result)(rawInput)

	assert.Nil(t, err)
	assert.Equal(t, float64(12.12), result)
}

func TestUtilInt64Marshaler(t *testing.T) {
	var input int64 = 12

	result, err := Int64Marshaler(input)()

	assert.Nil(t, err)
	assert.Equal(t, int64(12), result)
}

func TestUtilInt64UnmarshalerFromFloat64(t *testing.T) {
	var rawInput float64 = 12
	var result int64
	err := Int64Unmarshaler(&result)(rawInput)

	assert.Nil(t, err)
	assert.Equal(t, int64(12), result)
}

func TestUtilInt64UnmarshalerFromInt64(t *testing.T) {
	var rawInput int64 = 12
	var result int64
	err := Int64Unmarshaler(&result)(rawInput)

	assert.Nil(t, err)
	assert.Equal(t, int64(12), result)
}

func TestUtilStringMarshaler(t *testing.T) {
	var input string = "rgrert"

	result, err := StringMarshaler(input)()

	assert.Nil(t, err)
	assert.Equal(t, "rgrert", result)
}

func TestUtilStringUnmarshaler(t *testing.T) {
	var rawInput string = "aaa"
	var result string
	err := StringUnmarshaler(&result)(rawInput)

	assert.Nil(t, err)
	assert.Equal(t, "aaa", result)
}

func TestUtilBooleanMarshaler(t *testing.T) {
	var input bool = true

	result, err := BooleanMarshaler(input)()

	assert.Nil(t, err)
	assert.Equal(t, true, result)
}

func TestUtilBooleanUnmarshaler(t *testing.T) {
	var rawInput bool = true
	var result bool = false
	err := BooleanUnmarshaler(&result)(rawInput)

	assert.Nil(t, err)
	assert.Equal(t, true, result)
}

func TestUtilUnionTypeUnmarshaler(t *testing.T) {
	rawInput := map[string]interface{}{
		"type":              "type1",
		"payload_for_type1": "some value",
	}
	var result string
	err := UnionTypeUnmarshaler(func(typeStr string) unmarshaler {
		if typeStr == "type1" {
			return func(raw interface{}) error {
				assert.IsType(t, map[string]interface{}{}, raw)
				assert.IsType(t, "", raw.(mRaw)["type"])
				result = raw.(mRaw)["payload_for_type1"].(string)
				return nil
			}
		}
		require.Fail(t, "should not be reached")
		return nil
	})(rawInput)

	assert.Nil(t, err)
	assert.Equal(t, "some value", result)
}

func TestUtilEnumMarshaler(t *testing.T) {
	type enumType int64
	var enumVal1 enumType = 1
	var enumVal2 enumType = 5

	mapping := map[enumType]string{
		enumVal1: "val1",
		enumVal2: "val2",
	}

	var input enumType = enumVal2

	result, err := EnumMarshaler(input, mapping)()

	assert.Nil(t, err)
	assert.Equal(t, "val2", result)
}

func TestUtilEnumUnmarshaler(t *testing.T) {
	type enumType int64
	var enumVal1 enumType = 1
	var enumVal2 enumType = 5

	mapping := map[string]enumType{
		"val1": enumVal1,
		"val2": enumVal2,
	}
	var unmarshaledValue enumType
	err := EnumUnmarshaler(&unmarshaledValue, mapping)("val2")

	assert.Nil(t, err)
	assert.Equal(t, enumVal2, unmarshaledValue)
}

func TestUtilPtrMarshaler(t *testing.T) {
	var input int64 = 123
	result, err := PtrMarshaler(&input, Int64Marshaler)()

	assert.Nil(t, err)
	assert.Equal(t, int64(123), result)
}

func TestUtilPtrUnamrshaler(t *testing.T) {
	var unmarshaled *int64
	err := PtrUnmarshaler(&unmarshaled, Int64Unmarshaler)(int64(123))

	require.Nil(t, err)
	require.NotNil(t, unmarshaled)
	assert.Equal(t, int64(123), *unmarshaled)
}
