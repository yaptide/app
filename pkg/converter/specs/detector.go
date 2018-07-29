package specs

// DetectorID ...
type DetectorID int64

// Detector describes where and what values are scored during simulation.
type Detector struct {
	ID             DetectorID
	Geometry       DetectorGeometry
	ScoredParticle Particle
	Scoring        DetectorScoring
}

// DetectorGeometry ...
type DetectorGeometry interface {
	isDetectorGeometry() bool
}

func (d DetectorGeomap) isDetectorGeometry() bool {
	return true
}

func (d DetectorZones) isDetectorGeometry() bool {
	return true
}

func (d DetectorCylinder) isDetectorGeometry() bool {
	return true
}

func (d DetectorMesh) isDetectorGeometry() bool {
	return true
}

func (d DetectorPlane) isDetectorGeometry() bool {
	return true
}
