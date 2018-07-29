package web

import (
	"context"

	model "github.com/yaptide/yaptide/model/simulation"
)

func (h *handler) getServerConfiguration(ctx context.Context) (interface{}, error) {
	return struct {
		PredefinedMaterials []model.PredefinedMaterialRecord `json:"predefinedMaterials"`
		Isotopes            []model.IsotopeRecord            `json:"isotopes"`
		ParticleTypes       []model.PredefinedParticleRecord `json:"particles"`
		ScoringTypes        []model.ScoringTypeRecord        `json:"scoringTypes"`
	}{
		PredefinedMaterials: model.PredefinedMaterialsList,
		Isotopes:            model.IsotopesList,
		ParticleTypes:       model.PredefinedParticlesList,
		ScoringTypes:        model.ScoringTypesList,
	}, nil
}
