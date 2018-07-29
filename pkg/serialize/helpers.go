package serialize

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
	"runtime/debug"

	"github.com/davecgh/go-spew/spew"
	"github.com/yaptide/yaptide/pkg/converter/errors"
)

type mRaw = map[string]interface{}
type aRaw = []interface{}
type mErr = errors.MErr
type aErr = errors.AErr

type marshaler = func() (interface{}, error)
type unmarshaler = func(raw interface{}) error
type unmarshalTarget = interface{}

type fieldMarshaler = func(string, marshaler)
type fieldUnmarshaler = func(string, unmarshaler)

// StructMarshaler is util function that creates marshaler for struct.
//
//  val value ExampleStruct
//  marshaler := StructMarshaler(func(m fieldMarshaler) {
//    m("jsonField1", FieldMarshaler(value.Field1))
//    m("jsonField2", AnotherFieldMarshaler(value.Field2))
//  })
//
func StructMarshaler(marshal func(fieldMarshaler)) marshaler {
	return func() (interface{}, error) {
		errors := mErr{}
		result := map[string]interface{}{}
		marshal(func(field string, m marshaler) {
			value, err := m()
			result[field] = value
			if err != nil {
				errors[field] = err
			}
		})
		if len(errors) != 0 {
			return nil, errors
		}
		return result, nil
	}
}

// StructUnmarshaler is util function that creates unmarshaler for struct.
//
//  val value ExampleStruct
//  unmarshaler := StructUnmarshaler(func(u fieldUnmarshaler) {
//    u("jsonField1", FieldUnmarshaler(&value.Field1))
//    u("jsonField2", AnotherFieldUnmarshaler(&value.Field2))
//  })(raw)
//
func StructUnmarshaler(unmarshal func(fieldUnmarshaler)) unmarshaler {
	return func(raw interface{}) error {
		errors := mErr{}

		rawMap, isRawMap := raw.(map[string]interface{})
		if !isRawMap {
			return fmt.Errorf("expected map")
		}

		unmarshal(func(field string, u unmarshaler) {
			fieldValue, fieldExists := rawMap[field]
			if !fieldExists {
				return
			}
			unmarshalErr := u(fieldValue)
			if unmarshalErr != nil {
				errors[field] = unmarshalErr
			}
		})
		if len(errors) != 0 {
			return errors
		}
		return nil
	}
}

func ListMarshaler(
	list interface{}, marshalElement interface{},
) marshaler {
	return func() (interface{}, error) {
		listValue := reflect.ValueOf(list)
		if listValue.Kind() != reflect.Slice {
			return nil, fmt.Errorf("expected slice")
		}

		marshalElementFunType := reflect.TypeOf(marshalElement)
		if marshalElementFunType.Kind() != reflect.Func ||
			marshalElementFunType.NumIn() != 1 ||
			marshalElementFunType.NumOut() != 1 {
			panic("invalid function signature")
		}

		inputType := marshalElementFunType.In(0)
		outputType := marshalElementFunType.Out(0)
		if outputType != reflect.TypeOf((*marshaler)(nil)).Elem() {
			panic("wrong return type")
		}

		marshalElementFun := reflect.ValueOf(marshalElement)
		result := make([]interface{}, listValue.Len())
		errors := make(aErr, listValue.Len())
		hasErrors := false
		for i := 0; i < listValue.Len(); i++ {
			listElement := listValue.Index(i)
			if listElement.Type() != inputType {
				panicMarshaler("wrong type of element for marshaler constructor", list)
			}
			marshalerValueList := marshalElementFun.Call([]reflect.Value{
				listElement,
			})
			marshalerFunc := marshalerValueList[0].Interface().(marshaler)

			marshaledElement, elementErr := marshalerFunc()
			if elementErr != nil {
				errors[i] = elementErr
				hasErrors = true
			}
			result[i] = marshaledElement
		}
		if hasErrors {
			return result, errors
		}
		return result, nil
	}
}

func ListUnmarshaler(
	unpack interface{}, unmarshalElement interface{},
) unmarshaler {
	unpackValue := reflect.ValueOf(unpack)
	if unpackValue.Kind() != reflect.Ptr {
		panic("expected pointer to slice")
	}
	targetValue := unpackValue.Elem()
	if targetValue.Kind() != reflect.Slice {
		panic("expected pointer to slice")
	}

	return func(raw interface{}) error {
		rawSlice, isRawSlice := raw.([]interface{})
		if !isRawSlice {
			return fmt.Errorf("expected list of elements")
		}
		if targetValue.IsNil() {
			sliceType := reflect.SliceOf(targetValue.Type().Elem())
			slice := reflect.MakeSlice(sliceType, len(rawSlice), len(rawSlice))
			targetValue.Set(slice)
		}

		unmarshalElementFunType := reflect.TypeOf(unmarshalElement)
		if unmarshalElementFunType.Kind() != reflect.Func ||
			unmarshalElementFunType.NumIn() != 1 ||
			unmarshalElementFunType.NumOut() != 1 {
			panic("invalid function signature")
		}

		inputType := unmarshalElementFunType.In(0)
		outputType := unmarshalElementFunType.Out(0)
		if outputType != reflect.TypeOf((*unmarshaler)(nil)).Elem() {
			panic("wrong return type")
		}

		unmarshalElementFun := reflect.ValueOf(unmarshalElement)
		errors := make(aErr, len(rawSlice))
		hasErrors := false
		for i, element := range rawSlice {
			unpackElementTarget := reflect.New(inputType.Elem())

			unmarshalerValueList := unmarshalElementFun.Call([]reflect.Value{
				unpackElementTarget,
			})
			unmarshalFunc := unmarshalerValueList[0].Interface().(unmarshaler)

			err := unmarshalFunc(element)
			if err != nil {
				hasErrors = true
				errors[i] = err
			}

			targetValue.Index(i).Set(unpackElementTarget.Elem())
		}
		if hasErrors {
			return errors
		}
		return nil
	}
}

func Float64Marshaler(arg float64) marshaler {
	return func() (interface{}, error) {
		return arg, nil
	}
}

func Float64Unmarshaler(arg *float64) unmarshaler {
	return func(raw interface{}) error {
		switch t := raw.(type) {
		case float64:
			*arg = t
		}
		return nil
	}
}

func Int64Marshaler(arg int64) marshaler {
	return func() (interface{}, error) {
		return arg, nil
	}
}

func Int64Unmarshaler(arg *int64) unmarshaler {
	return func(raw interface{}) error {
		switch t := raw.(type) {
		case float64:
			if math.Trunc(t) != t {
				return fmt.Errorf("integer expected")
			}
			*arg = int64(t)
		case int64:
			*arg = t
		}
		return nil
	}
}

func StringMarshaler(arg string) marshaler {
	return func() (interface{}, error) {
		return arg, nil
	}
}

func StringUnmarshaler(arg *string) unmarshaler {
	return func(raw interface{}) error {
		switch t := raw.(type) {
		case string:
			*arg = t
		case []byte:
			*arg = string(t)
		default:
			return fmt.Errorf("string is required")
		}
		return nil
	}
}

func BooleanMarshaler(arg bool) marshaler {
	return func() (interface{}, error) {
		return arg, nil
	}
}

func BooleanUnmarshaler(arg *bool) unmarshaler {
	return func(raw interface{}) error {
		switch t := raw.(type) {
		case bool:
			*arg = t
		default:
			return fmt.Errorf("bool value is required")
		}
		return nil
	}
}

// UnionTypeUnmarshaler is util function that creates umarshaler for union
// types represented by map with key "type" where type is providing
// identification for different types.
//
//  var unmarshaledString string
//  var unmarshaledInt int64
//  unmarshaler := UnionTypeUnmarshaler(func(unionType string) unmarshaler {
//	  if unionType == "type1" {
//      return StringUnmarshaler(&unmarshaledString)
//    } else if unionType == "type2" {
//      return Int64Unmarshaler(&unmarshaledInt)
//    }
//  })
func UnionTypeUnmarshaler(
	unmarshal func(t string) unmarshaler,
) unmarshaler {
	return func(raw interface{}) error {
		if raw == nil {
			return nil
		}
		rawMap, isMap := raw.(map[string]interface{})
		if !isMap {
			return fmt.Errorf("expected map")
		}
		typeObject, exists := rawMap["type"]
		if !exists {
			return fmt.Errorf("field type is required")
		}
		typeStr, isString := typeObject.(string)
		if !isString {
			return fmt.Errorf("field type need to be a string")
		}
		return unmarshal(typeStr)(raw)
	}
}

// EnumMarshaler is util function that creates marshaler for set of enum
// values. Value is marshaled based on mapping provided to EnumMarshaler
// function.
// WARNING: enum value can't be zero value fr underlying type.
//
//  mapping := map[enumType]string{
//		value1: "value1",
//		value2: "value2",
//  }
//  marshaler := EnumMarshaler(value1, mapping)
//
func EnumMarshaler(
	value interface{}, mapping interface{},
) marshaler {
	mappingValue := reflect.ValueOf(mapping)
	if mappingValue.Kind() != reflect.Map {
		panic("expected map")
	}
	mappingType := mappingValue.Type()
	valueType := reflect.TypeOf(value)
	if mappingType.Key() != valueType {
		panic("mapping key type has to be the same as passed value")
	}

	return func() (interface{}, error) {
		enumValue := mappingValue.MapIndex(reflect.ValueOf(value))
		if !enumValue.IsValid() {
			return nil, fmt.Errorf("unknown enum value")
		}

		return enumValue.Interface(), nil
	}
}

// EnumUnmarshaler is util function that creates unmarshaler for set of enum
// values. Value is unmarshaled based on mapping provided to EnumUnmarshaler
// function.
// WARNING: enum value can't be zero value fr underlying type.
//
//  mapping := map[string]enumType{
//		"value1": value1,
//		"value2": value2,
//  }
//  var unmarshaledValue enumType
//  marshaler := EnumUnmarshaler(&unmarshaledValue, mapping)
//
func EnumUnmarshaler(
	unpack interface{}, mapping interface{},
) unmarshaler {
	mappingValue := reflect.ValueOf(mapping)
	if mappingValue.Kind() != reflect.Map {
		panic("expected map")
	}
	mappingType := mappingValue.Type()
	unpackValue := reflect.ValueOf(unpack)
	if unpackValue.Kind() != reflect.Ptr {
		panic("unpack target need to be a pointer")
	}
	targetValue := unpackValue.Elem()

	if mappingType.Elem() != targetValue.Type() {
		panic("mapping value type has to be the same as value pointed by passed ptr")
	}
	mappingKeyType := mappingType.Key()

	return func(raw interface{}) error {
		rawValue := reflect.ValueOf(raw)
		rawValueType := rawValue.Type()

		if !rawValueType.ConvertibleTo(mappingKeyType) {
			return fmt.Errorf(
				"can't convert %v to %v",
				rawValue.Type(), mappingKeyType,
			)
		}
		if mappingKeyType != rawValueType {
			rawValue = rawValue.Convert(mappingKeyType)
		}

		enumValue := mappingValue.MapIndex(rawValue)
		if enumValue == reflect.Zero(mappingType.Elem()) {
			return fmt.Errorf("unknown enum value")
		}
		targetValue.Set(enumValue)
		return nil
	}
}

func PtrMarshaler(
	ptr interface{}, newMarshaler interface{},
) marshaler {
	ptrValue := reflect.ValueOf(ptr)
	if ptrValue.Kind() != reflect.Ptr {
		panic("this function need to have signature PtrMarshaler(ptr *T, func(T) marshaler)")
	}

	newMarshalerValue := reflect.ValueOf(newMarshaler)
	newMarshalerType := newMarshalerValue.Type()
	if newMarshalerValue.Kind() != reflect.Func ||
		newMarshalerType.NumIn() != 1 ||
		newMarshalerType.NumOut() != 1 {
		panic("this function need to have signature PtrMarshaler(ptr *T, func(T) marshaler)")
	}

	inputType := newMarshalerType.In(0)
	outputType := newMarshalerType.Out(0)
	if outputType != reflect.TypeOf((*marshaler)(nil)).Elem() {
		panic("wrong return type")
	}

	if inputType != ptrValue.Type().Elem() {
		panic("first arg and pointed should be the same type")
	}

	return func() (interface{}, error) {
		if ptrValue.IsNil() {
			return nil, nil
		}
		marshalerFunSlice := newMarshalerValue.Call([]reflect.Value{
			ptrValue.Elem(),
		})
		return marshalerFunSlice[0].Interface().(marshaler)()
	}
}

func PtrUnmarshaler(
	ptr interface{}, newUnmarshaler interface{},
) unmarshaler {
	ptrToPtrValue := reflect.ValueOf(ptr)
	if ptrToPtrValue.Kind() != reflect.Ptr ||
		ptrToPtrValue.Elem().Kind() != reflect.Ptr {
		panic("expected signature func(p **T, func(u *T) unmarshaler)")
	}

	newUnmarshalerValue := reflect.ValueOf(newUnmarshaler)
	newUnmarshalerType := newUnmarshalerValue.Type()
	if newUnmarshalerType.Kind() != reflect.Func ||
		newUnmarshalerType.NumIn() != 1 ||
		newUnmarshalerType.NumOut() != 1 {
		panic("expected signature func(p **T, func(u *T) unmarshaler)")
	}

	inputType := newUnmarshalerType.In(0)
	outputType := newUnmarshalerType.Out(0)
	if outputType != reflect.TypeOf((*unmarshaler)(nil)).Elem() {
		panic("wrong return type")
	}

	if inputType != ptrToPtrValue.Elem().Type() {
		panic("first arg and pointed should be the same type")
	}
	ptrToPtrValue.Elem().Set(reflect.New(ptrToPtrValue.Type().Elem().Elem()))

	return func(raw interface{}) error {
		unmarshalerFunSlice := newUnmarshalerValue.Call([]reflect.Value{
			ptrToPtrValue.Elem(),
		})
		return unmarshalerFunSlice[0].Interface().(unmarshaler)(raw)
	}
}

func panicMarshaler(msg string, element interface{}) {
	funPC, _, _, _ := runtime.Caller(1)
	funcName := runtime.FuncForPC(funPC).Name()
	panicMsg := fmt.Sprintf(
		"\n%s %s\n%s: %s\n%s\n%s",
		"Error while constructing marshaler", funcName,
		"Failed with error", msg,
		"Marshaled value", spew.Sdump(element),
	)
	debug.PrintStack()
	panic(panicMsg)
}
