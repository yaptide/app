package web

import (
	"context"

	"github.com/yaptide/yaptide/model"
	"gopkg.in/mgo.v2/bson"
)

func (h *handler) getSimulationResults(
	ctx context.Context,
) (*model.SimulationResultsDB, error) {
	a := extractActionContext(ctx)
	resultsID := extractSimulationResultID(ctx)

	result, resultErr := h.Resolver.SimulationResultsGet(a, resultsID)
	if resultErr != nil {
		return nil, resultErr
	}

	return result, nil
}

func (h *handler) getSimulationSpecs(
	ctx context.Context,
) (*model.SimulationSpecsDB, error) {
	a := extractActionContext(ctx)
	specsID := extractSimulationSpecsID(ctx)

	specs, specsErr := h.Resolver.SimulationSpecsGet(a, specsID)
	if specsErr != nil {
		return nil, specsErr
	}

	return specs, nil
}

func (h *handler) updateSimulationSpecs(
	ctx context.Context, input *model.SimulationSpecs,
) (*model.SimulationSpecsDB, error) {
	a := extractActionContext(ctx)
	specsID := extractSimulationSpecsID(ctx)

	if err := h.Resolver.SimulationSpecsUpdate(a, specsID, input); err != nil {
		return nil, err
	}

	specs, getErr := h.Resolver.SimulationSpecsGet(a, specsID)
	if getErr != nil {
		return nil, getErr
	}

	return specs, nil
}

func (h *handler) runSimulationHandler(
	ctx context.Context,
	args *struct {
		ProjectID bson.ObjectId `json:"projectId"`
		VersionID int           `json:"versionId"`
	},
) error {
	userID := extractUserID(ctx)
	return h.simulationHandler.HandleSimulation(args.ProjectID, args.VersionID, userID)
}
