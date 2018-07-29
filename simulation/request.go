package simulation

import (
	model "github.com/yaptide/yaptide/model"
)

type request interface {
	ConvertModel(specs *model.SimulationSpecs) error
	StartSimulation() error
}
