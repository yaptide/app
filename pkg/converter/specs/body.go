package specs

// BodyID is key type in Body map.
type BodyID int64

// Body store Geometry interface described by ID and Name.
type Body struct {
	ID       BodyID       `json:"id"`
	Geometry BodyGeometry `json:"geometry"`
}

// BodyGeometry ...
type BodyGeometry interface {
	isBodyType() bool
}

func (b BodySphere) isBodyType() bool {
	return true
}

func (b BodyCuboid) isBodyType() bool {
	return true
}

func (b BodyCylinder) isBodyType() bool {
	return true
}
