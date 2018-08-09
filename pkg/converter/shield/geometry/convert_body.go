package geometry

import (
	"fmt"
	"sort"

	"github.com/yaptide/yaptide/pkg/converter/geometry"
	"github.com/yaptide/yaptide/pkg/converter/specs"
)

// ShieldBodyID ...
type ShieldBodyID int64

// Body represent specs.Body,
type Body struct {
	ID         ShieldBodyID
	Identifier string
	Arguments  []float64
}

func convertBodies(
	bodyMap map[specs.BodyID]specs.Body,
) ([]Body, map[specs.BodyID]ShieldBodyID, error) {
	result := []Body{}
	bodyIDToShield := map[specs.BodyID]ShieldBodyID{}

	bodyIds := []specs.BodyID{}
	for k := range bodyMap {
		bodyIds = append(bodyIds, k)
	}
	sort.SliceStable(bodyIds, func(i, j int) bool { return bodyIds[i] < bodyIds[j] })

	for i, id := range bodyIds {
		nextShieldID := ShieldBodyID(i + 1)
		bodyIDToShield[id] = nextShieldID

		body, err := convertBody(bodyMap[id])
		if err != nil {
			return nil, nil, err
		}
		body.ID = nextShieldID
		result = append(result, body)
	}
	return result, bodyIDToShield, nil
}

func appendBlackholeBody(bodies []Body) ([]Body, ShieldBodyID, error) {
	newID := bodies[len(bodies)-1].ID + 1

	blackholeBody, err := convertCuboid(specs.BodyCuboid{
		Center: geometry.Point{
			X: 0.0,
			Y: 0.0,
			Z: 0.0,
		},
		Size: geometry.Vec3D{
			X: 500.0,
			Y: 500.0,
			Z: 500.0,
		},
	})

	if err != nil {
		return nil, 0, err
	}

	blackholeBody.ID = newID
	return append(bodies, blackholeBody), newID, nil

}

func convertBody(b specs.Body) (Body, error) {
	switch g := b.Geometry.(type) {
	case specs.BodySphere:
		return convertSphere(g)
	case specs.BodyCuboid:
		return convertCuboid(g)
	case specs.BodyCylinder:
		return convertCylinder(g)

	default:
		return Body{}, fmt.Errorf("geometry type %T serializing not implemented", b.Geometry)
	}
}

func convertSphere(sphere specs.BodySphere) (Body, error) {
	if sphere.Radius <= 0.0 {
		return Body{}, fmt.Errorf("sphere radius cannot be <= 0.0")
	}

	return Body{
		Identifier: "SPH",
		Arguments:  []float64{sphere.Center.X, sphere.Center.Y, sphere.Center.Z, sphere.Radius},
	}, nil
}

func convertCuboid(cuboid specs.BodyCuboid) (Body, error) {
	for axis, size := range map[string]float64{
		"x": cuboid.Size.X,
		"y": cuboid.Size.Y,
		"z": cuboid.Size.Z,
	} {
		if size <= 0.0 {
			return Body{}, fmt.Errorf("cuboid size in %s axis cannot be <= 0.0", axis)
		}
	}

	minX, maxX := geometry.CenterAndSizeToMinAndMax(cuboid.Center.X, cuboid.Size.X)
	minY, maxY := geometry.CenterAndSizeToMinAndMax(cuboid.Center.Y, cuboid.Size.Y)
	minZ, maxZ := geometry.CenterAndSizeToMinAndMax(cuboid.Center.Z, cuboid.Size.Z)

	return Body{
		Identifier: "RPP",
		Arguments:  []float64{minX, maxX, minY, maxY, minZ, maxZ},
	}, nil
}

func convertCylinder(cylinder specs.BodyCylinder) (Body, error) {
	if cylinder.Height <= 0.0 {
		return Body{}, fmt.Errorf("cylinder height cannot be <= 0.0")
	}

	if cylinder.Radius <= 0.0 {
		return Body{}, fmt.Errorf("cylinder radius cannot be <= 0.0")
	}

	// TODO: support cylinders, which vector from the center to the opposite end of the
	// cylinder are not parallel to [0, 1, 0].
	return Body{
		Identifier: "RCC",
		Arguments: []float64{
			cylinder.Center.X,
			cylinder.Center.Y,
			cylinder.Center.Z,
			0,
			cylinder.Height, 0,
			cylinder.Radius,
		},
	}, nil
}
