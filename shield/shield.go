// Package shield implements converters between model and input/output for SHIELD-HIT12A library.
package shield

import (
	"github.com/yaptide/converter"
	"github.com/yaptide/converter/shield/context"
	"github.com/yaptide/converter/shield/results"
	"github.com/yaptide/converter/shield/setup"
)

func ConvertSetup(simulationSetup converter.Setup) (setup.RawShieldSetup, SerializationContext, error) {
	return setup.Convert(simulationSetup)
}

func ParseResults(files map[string]string, simulationContext *SerializationContext) (*converter.Result, error) {
	return results.ParseResults(files, simulationContext)
}

type SerializationContext = context.SerializationContext
type RawShieldSetup = setup.RawShieldSetup
