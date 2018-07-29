package model

import (
	"github.com/yaptide/yaptide/model/simulation"
	"github.com/yaptide/yaptide/pkg/converter"
	"gopkg.in/mgo.v2/bson"
)

type SimulationSpecsDB = simulation.SpecsDB
type SimulationResultsDB = simulation.ResultsDB

type SimulationSpecs = simulation.Specs
type SimulationResults = converter.Result

// InitialSimulationSpcs
func InitialSimulationSpecs(userID bson.ObjectId) *SimulationSpecsDB {
	return &SimulationSpecsDB{
		ID:     bson.NewObjectId(),
		UserID: userID,
	}
}

// InitialSimulationResult ...
func InitialSimulationResult(userID bson.ObjectId) *SimulationResultsDB {
	return &SimulationResultsDB{
		ID:     bson.NewObjectId(),
		UserID: userID,
	}
}
