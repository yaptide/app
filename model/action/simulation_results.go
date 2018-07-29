package action

import (
	"github.com/yaptide/yaptide/errors"
	"github.com/yaptide/yaptide/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// SimulationResultsGet ...
func (r *Resolver) SimulationResultsGet(
	ctx *context, resultsID bson.ObjectId,
) (*model.SimulationResultsDB, error) {
	result := &model.SimulationResultsDB{}
	getErr := ctx.db.SimulationResult().Find(M{
		"_id":    resultsID,
		"userId": ctx.userID,
	}).One(result)
	if getErr == mgo.ErrNotFound {
		return nil, errors.ErrNotFound
	}
	if getErr != nil {
		return nil, errors.ErrInternalServerError
	}
	if result.UserID != ctx.userID {
		return nil, errors.ErrUnauthorized
	}
	return result, nil
}

// SimulationResultsCreateInitial ...
func (r *Resolver) SimulationResultsCreateInitial(
	ctx *context,
) (*model.SimulationResultsDB, error) {
	result := model.InitialSimulationResult(ctx.userID)
	insertErr := ctx.db.SimulationResult().Insert(result)
	if insertErr != nil {
		return nil, errors.ErrInternalServerError
	}
	return result, nil
}
