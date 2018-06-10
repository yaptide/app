package serialize

import (
	conf "github.com/yaptide/yaptide/config"
	"github.com/yaptide/yaptide/pkg/converter/errors"
)

type mRaw = map[string]interface{}
type aRaw = []interface{}
type mErr = errors.MErr
type aErr = errors.AErr

var log = conf.NamedLogger("converter/serialize")
