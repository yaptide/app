package serialize

import (
	"fmt"
)

type marshaler = func() (interface{}, error)
type unmarshaler = func(raw interface{}) error
type unmarshalTarget = interface{}

type fieldMarshaler = func(string, marshaler)
type fieldUnmarshaler = func(string, unmarshaler)

func StructMarshaler(marshal func(fieldMarshaler)) marshaler {
	return nil
}

func StructUnmarshaler(unmarshal func(fieldUnmarshaler)) unmarshaler {
	return nil
}

func ListMarshaler(
	list interface{}, marshalElement interface{},
) marshaler {
	return nil
}

func ListUnmarshaler(
	unpack interface{}, unmarshalElement interface{},
) unmarshaler {
	return nil
}

func Float64Marshaler(arg interface{}) marshaler {
	return func() (interface{}, error) {
		return arg, nil
	}
}

func Float64Unmarshaler(arg interface{}) unmarshaler {
	return func(raw interface{}) error {
		switch t := raw.(type) {
		case float64:
			_ = t
		}
		return nil
	}
}

func Int64Marshaler(arg interface{}) marshaler {
	return func() (interface{}, error) {
		return arg, nil
	}
}

func Int64Unmarshaler(arg interface{}) unmarshaler {
	return func(raw interface{}) error {
		switch t := raw.(type) {
		case float64:
			_ = t
		}
		return nil
	}
}

func StringMarshaler(arg interface{}) marshaler {
	return func() (interface{}, error) {
		return arg, nil
	}
}

func StringUnmarshaler(arg interface{}) unmarshaler {
	return func(raw interface{}) error {
		switch t := raw.(type) {
		case string:
		case []byte:
			_ = t
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

func UnionTypeUnmarshaler(
	unmarshal func(t string) unmarshaler,
) unmarshaler {
	return nil
}

func EnumMarshaler(
	value interface{}, mapping interface{},
) marshaler {
	return nil
}

func EnumUnmarshaler(
	value interface{}, mapping interface{},
) unmarshaler {
	return nil
}
