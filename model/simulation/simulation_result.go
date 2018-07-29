package simulation

import (
	"github.com/yaptide/yaptide/pkg/converter"
	"gopkg.in/mgo.v2/bson"
)

// SimulationResult ...
type ResultsDB struct {
	ID               bson.ObjectId `json:"id" bson:"_id"`
	UserID           bson.ObjectId `json:"userId" bson:"userId"`
	converter.Result `bson:",inline"`
}
