package action

import (
	"github.com/yaptide/yaptide/errors"
	"github.com/yaptide/yaptide/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// SimulationSpecsGet ...
func (r *Resolver) SimulationSpecsGet(
	ctx *context, specsID bson.ObjectId,
) (*model.SimulationSpecsDB, error) {
	specs := model.SimulationSpecsDB{}
	log.Error(specsID, ctx.userID)
	getErr := ctx.db.SimulationSpecs().Find(bson.M{
		"_id":    specsID,
		"userId": ctx.userID,
	}).One(&specs)
	if getErr == mgo.ErrNotFound {
		log.Warn(getErr.Error())
		return nil, errors.ErrNotFound
	}
	if getErr != nil {
		log.Warn(getErr.Error())
		return nil, errors.ErrInternalServerError
	}
	if specs.UserID != ctx.userID {
		return nil, errors.ErrUnauthorized
	}
	return &specs, nil
}

// SimulationSpecsCreateInitial creates default starting simulation and
// inserts it into db.
func (r *Resolver) SimulationSpecsCreateInitial(
	ctx *context,
) (*model.SimulationSpecsDB, error) {
	specs := model.InitialSimulationSpecs(ctx.userID)
	insertErr := ctx.db.SimulationSpecs().Insert(specs)
	if insertErr != nil {
		log.Warnf(
			"SimulationSpecs initial insert failed for specs %s with error[%s]",
			specs.ID.Hex(), insertErr.Error(),
		)
		return nil, errors.ErrInternalServerError
	}
	return specs, nil
}

// SimulationSpecsCreateFrom creates copy of simulation specs selected by id
// and inserts it into db.
func (r *Resolver) SimulationSpecsCreateFrom(
	ctx *context, specsID bson.ObjectId,
) (*model.SimulationSpecsDB, error) {
	specs, getErr := r.SimulationSpecsGet(ctx, specsID)
	if getErr != nil {
		log.Warnf(
			"SimulationSpecs get failed for specs %s",
			specsID.Hex(),
		)
		return nil, getErr
	}
	specs.ID = bson.NewObjectId()
	specs.UserID = ctx.userID
	insertErr := ctx.db.SimulationSpecs().Insert(specs)
	if insertErr != nil {
		log.Warnf(
			"SimulationSpecs insert failed for specs %s with error[%s]",
			specsID.Hex(), insertErr.Error(),
		)
		return nil, errors.ErrInternalServerError
	}
	return specs, nil
}

// SimulationSpecsUpdate relaces entire specs object inside db.
func (r *Resolver) SimulationSpecsUpdate(
	ctx *context, specsID bson.ObjectId, input *model.SimulationSpecs,
) error {
	updateErr := ctx.db.SimulationSpecs().Update(M{
		"_id":    specsID,
		"userId": ctx.userID,
	}, M{
		"$set": input,
	})
	if updateErr != nil {
		return errors.ErrInternalServerError
	}
	return nil
}
