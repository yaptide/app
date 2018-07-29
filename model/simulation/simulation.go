package simulation

import (
	conf "github.com/yaptide/yaptide/config"
	"github.com/yaptide/yaptide/pkg/converter/errors"
)

type mRaw = map[string]interface{}
type aRaw = []interface{}
type mErr = errors.MErr
type aErr = errors.AErr

type marshaler = func() (interface{}, error)
type unmarshaler = func(raw interface{}) error

type fieldMarshaler = func(string, marshaler)
type fieldUnmarshaler = func(string, unmarshaler)

var log = conf.NamedLogger("converter/serialize")
