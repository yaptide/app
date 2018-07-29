package geometry

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yaptide/yaptide/pkg/converter/specs"
)

func TestOetFromZoneDescription(t *testing.T) {
	baseBodyID := ShieldBodyID(1)
	operations := []operation{
		operation{BodyID: 2, Type: specs.Intersect},
		operation{BodyID: 3, Type: specs.Union},
		operation{BodyID: 4, Type: specs.Subtract},
	}

	expected := createOetBinaryExpression(intersection,
		createOetBinaryExpression(
			union,
			createOetBinaryExpression(
				intersection,
				createOetValue(1, Plus),
				createOetValue(2, Plus),
				Plus,
			),
			createOetValue(3, Plus),
			Plus,
		),
		createOetValue(4, Minus),
		Plus)

	actual := oetFromZoneDescription(baseBodyID, operations)

	assert.Equal(t, expected, actual)
}

func TestOetUnion(t *testing.T) {
	oets := []*oet{
		createOetValue(1, Plus),
		createOetValue(2, Minus),
		createOetValue(3, Minus),
	}

	expected := createOetBinaryExpression(
		union,
		createOetBinaryExpression(
			union,
			createOetValue(1, Plus),
			createOetValue(2, Minus),
			Plus,
		),
		createOetValue(3, Minus),
		Plus,
	)

	actual := oetUnion(oets...)

	assert.Equal(t, expected, actual)
}

func TestOetSubstract(t *testing.T) {
	left := createOetValue(1, Plus)
	right := createOetBinaryExpression(
		union,
		createOetBinaryExpression(
			union,
			createOetValue(1, Plus),
			createOetValue(2, Minus),
			Plus,
		),
		createOetValue(3, Minus),
		Plus,
	)

	expected := createOetBinaryExpression(
		intersection,
		createOetValue(1, Plus),
		createOetBinaryExpression(
			union,
			createOetBinaryExpression(
				union,
				createOetValue(1, Plus),
				createOetValue(2, Minus),
				Plus,
			),
			createOetValue(3, Minus),
			Minus,
		),
		Plus)

	actual := oetSubtract(left, right)

	assert.Equal(t, expected, actual)
}
