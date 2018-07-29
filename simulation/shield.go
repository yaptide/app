package simulation

import (
	"github.com/yaptide/yaptide/model"
	"github.com/yaptide/yaptide/pkg/converter/shield"
	"github.com/yaptide/yaptide/runner/file"
)

type shieldProcessor struct {
	shield.RawShieldSpecs
	serializationContext shield.SerializationContext
}

func (p *shieldProcessor) ConvertModel(simulationsSpecs *model.SimulationSpecs) error {
	convertedModel, serializationContext, convertErr := shield.Convert(
		simulationsSpecs.ConverterConfiguration(),
	)
	if convertErr != nil {
		return convertErr
	}
	p.RawShieldSpecs = convertedModel
	p.serializationContext = serializationContext
	return nil
}

func (p *shieldProcessor) HandleFileResults(
	simResults file.SimulationResults,
) (*model.SimulationResults, error) {
	parserOutput, parserErr := shield.ParseResults(
		simResults.Files,
		&p.serializationContext,
	)
	if parserErr != nil {
		return nil, parserErr
	}
	return parserOutput, nil
}
