package setup

// ZoneOperationType determines operation type.
// OperationTypes are based on mathematical operations on sets.
type ZoneOperationType int64

const (
	// Intersect operation: A ∩ B.
	Intersect ZoneOperationType = iota
	// Subtract operation: A \ B.
	Subtract
	// Union operation: A ∪ B.
	Union
)

// ZoneOperation determines construction of Zone.
type ZoneOperation struct {
	BodyID BodyID            `json:"bodyId"`
	Type   ZoneOperationType `json:"-"`
}

type rawOperation struct {
	BodyID BodyID `json:"bodyId"`
	Type   string `json:"type"`
}
