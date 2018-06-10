package serialize

import (
	"fmt"
	"github.com/yaptide/yaptide/pkg/converter/setup"
)

type StateOfMatter struct {
	*setup.StateOfMatter
}

var mapStateToJSON = map[setup.StateOfMatter]string{
	setup.UndefinedStateOfMatter: "",
	setup.Solid:                  "solid",
	setup.Gas:                    "gas",
	setup.Liquid:                 "liquid",
}

// reverse mapStateToJSON
var mapJSONToState = func() (mapping map[string]setup.StateOfMatter) {
	for key, value := range mapStateToJSON {
		mapping[value] = key
	}
	return mapping
}()

func (s StateOfMatter) Marshal() (interface{}, error) {
	val, ok := mapStateToJSON[*s.StateOfMatter]
	if !ok {
		log.Error("unknow state of matter %d", val)
		return "", fmt.Errorf("unknown state of matter")
	}
	return val, nil
}

func (s *StateOfMatter) Unmarshal(raw interface{}) error {
	rawStr, isStr := raw.(string)
	state, exists := mapJSONToState[rawStr]
	if !isStr || !exists {
		return fmt.Errorf("one of %v is required", mapStateToJSON)
	}
	*s.StateOfMatter = state
	return nil
}
