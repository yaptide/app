package specs

// DetectorScoring ...
type DetectorScoring interface {
	isDetectorScoring() bool
}

func (s PredefinedScoring) isDetectorScoring() bool {
	return true
}

func (s LetTypeScoring) isDetectorScoring() bool {
	return true
}

// PredefinedScoring ...
type PredefinedScoring string

// LetTypeScoring ...
type LetTypeScoring struct {
	Type     string
	Material MaterialID
}

// Validate ...
func (s PredefinedScoring) Validate() error {
	return nil
}

// Validate ...
func (s LetTypeScoring) Validate() error {
	return nil
}
