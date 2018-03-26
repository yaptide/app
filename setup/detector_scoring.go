package setup

import (
	"encoding/json"
	"fmt"

	"github.com/yaptide/converter/common"
)

// PredefinedScoring ...
type PredefinedScoring string

// LetTypeScoring ...
type LetTypeScoring struct {
	Type     string     `json:"type"`
	Material MaterialID `json:"material"`
}

// MarshalJSON json.Marshaller implementation.
func (g PredefinedScoring) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type string `json:"type"`
	}{
		Type: string(g),
	})
}

// MarshalJSON json.Marshaller implementation.
func (s LetTypeScoring) MarshalJSON() ([]byte, error) {
	type Alias LetTypeScoring
	return json.Marshal(struct {
		Alias
	}{
		Alias: (Alias)(s),
	})
}

func unmarshalScoringType(scoring json.RawMessage) (common.Particle, error) {
	var predefinedScoring struct {
		Type PredefinedScoring `json:"type"`
	}
	predefinedScoringErr := json.Unmarshal(scoring, &predefinedScoring)
	var letTypeScoring LetTypeScoring
	letTypeScoringErr := json.Unmarshal(scoring, &letTypeScoring)

	if letTypeScoringErr == nil && (letTypeScoring.Type == "letflu" ||
		letTypeScoring.Type == "dlet" ||
		letTypeScoring.Type == "tlet") {
		return letTypeScoring, nil
	}
	if predefinedScoringErr == nil {
		return predefinedScoring.Type, nil
	}
	return nil, fmt.Errorf("unknown scoring type")
}
