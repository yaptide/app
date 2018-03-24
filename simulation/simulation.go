// Package processor implements processing all simulation requests. Is responsible for serialization, starting simulation and processing results.
package simulation

import (
	"fmt"

	conf "github.com/yaptide/app/config"
	"github.com/yaptide/app/errors"
	"github.com/yaptide/app/model"
	"github.com/yaptide/app/model/action"
	"github.com/yaptide/app/model/mongo"
	"github.com/yaptide/app/runner/file"
	"gopkg.in/mgo.v2/bson"
)

var log = conf.NamedLogger("simulation_handler")

// Handler responsible for all steps of processing simulation.
type Handler struct {
	action                *action.Resolver
	db                    mongo.DB
	shieldFileLocalRunner *file.Runner
}

// NewProcessor constructor.
func NewHandler(action *action.Resolver, db mongo.DB) *Handler {
	processor := &Handler{
		action: action,
		db:     db,
		shieldFileLocalRunner: file.SetupShieldRunner(action.Config),
	}
	return processor
}

// HandleSimulation processes simulation.
func (p *Handler) HandleSimulation(projectID bson.ObjectId, versionID int, userID bson.ObjectId) error {
	version, versionErr := p.action.ProjectVersionGet(p.db, projectID, versionID, userID)
	if versionErr != nil {
		log.Warning("[SimulationProcessor] Unable to fetch version. Reason: %s", versionErr.Error())
		return versionErr
	}

	setup, setupErr := p.action.SimulationSetupGet(p.db, version.SetupID, userID)
	if setupErr != nil {
		return setupErr
	}

	if err := version.Settings.IsValid(); err != nil {
		return err
	}
	log.Debug("[SimulationProcessor] Start simulation request")

	request, requestErr := p.selectRequestFormSettings(version, projectID)
	if requestErr != nil {
		return requestErr
	}

	log.Debug("Start simulation request (serialization)")
	serializeErr := request.ConvertModel(setup)
	if serializeErr != nil {
		return serializeErr
	}

	log.Debug("[SimulationProcessor] Start simulation request (enqueue in runner)")
	startSimulationErr := request.StartSimulation()
	if startSimulationErr != nil {
		return startSimulationErr
	}

	return nil
}

func (h *Handler) selectRequestFormSettings(
	version *model.Version, projectID bson.ObjectId,
) (request, error) {
	switch version.Settings.ComputingLibrary {
	case model.ShieldLibrary:
		switch version.Settings.SimulationEngine {
		case model.LocalMachine:
			return &fileRequest{
				Runner:        h.shieldFileLocalRunner,
				fileProcessor: &shieldProcessor{},
				SimulationContext: h.action.NewSimulationContext(
					h.db,
					projectID,
					version.ID,
				),
			}, nil
		default:
			return nil, errors.ErrInternalServerError
		}
	case model.FlukaLibrary:
		return nil, errors.ErrNotImplemented
	default:
		return nil, fmt.Errorf("[SimulationProcessor] Invalid computing library")
	}
}
