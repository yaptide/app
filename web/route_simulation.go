package web

import (
	"context"

	"github.com/yaptide/app/model"
	"github.com/yaptide/converter"
	"gopkg.in/mgo.v2/bson"
)

func (h *handler) getSimulationResult(
	ctx context.Context,
) (*model.SimulationResult, error) {
	a := extractActionContext(ctx)
	resultID := extractSimualtionSetupId(ctx)

	result, resultErr := h.Resolver.SimulationResultGet(a, resultID)
	if resultErr != nil {
		return nil, resultErr
	}

	return result, nil
}

func (h *handler) getSimulationSetup(
	ctx context.Context,
) (*model.SimulationSetup, error) {
	a := extractActionContext(ctx)
	setupID := extractSimualtionSetupId(ctx)

	setup, setupErr := h.Resolver.SimulationSetupGet(a, setupID)
	if setupErr != nil {
		return nil, setupErr
	}

	return setup, nil
}

func (h *handler) updateSimulationSetup(
	input *converter.Setup,
	ctx context.Context,
) (*model.SimulationSetup, error) {
	a := extractActionContext(ctx)
	setupID := extractSimualtionSetupId(ctx)

	if err := h.Resolver.SimulationSetupUpdate(a, setupID, input); err != nil {
		return nil, err
	}

	setup, getErr := h.Resolver.SimulationSetupGet(a, setupID)
	if getErr != nil {
		return nil, getErr
	}

	return setup, nil
}

func (h *handler) runSimulationHandler(
	args *struct {
		ProjectID bson.ObjectId `json:"projectId"`
		VersionID int           `json:"versionId"`
	},
	ctx context.Context,
) error {
	userID := extractUserId(ctx)
	return h.simulationHandler.HandleSimulation(args.ProjectID, args.VersionID, userID)
}
