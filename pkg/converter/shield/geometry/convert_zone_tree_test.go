package geometry

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yaptide/yaptide/pkg/converter/shield/material"
	"github.com/yaptide/yaptide/pkg/converter/specs"
)

func TestConvertZonesToZoneTreeForest(t *testing.T) {
	type testCase struct {
		ZoneMap            map[specs.ZoneID]specs.Zone
		MaterialIDToShield map[specs.MaterialID]material.ShieldID
		BodyIDToShield     map[specs.BodyID]ShieldBodyID

		Expected      []*zoneTree
		ExpectedError error
	}

	check := func(t *testing.T, tc testCase) {
		t.Helper()

		actual, actualErr := convertZonesToZoneTreeForest(
			tc.ZoneMap, tc.MaterialIDToShield, tc.BodyIDToShield,
		)

		assert.Equal(t, tc.ExpectedError, actualErr)
		assert.Equal(t, tc.Expected, actual)
	}

	t.Run("SimpleOneZone", func(t *testing.T) {
		check(t, testCase{
			ZoneMap: createZoneMap(specs.Zone{
				ID:         specs.ZoneID(1),
				ParentID:   specs.ZoneID(0),
				BaseID:     specs.BodyID(1),
				MaterialID: specs.MaterialID(2),
				Construction: []specs.ZoneOperation{
					specs.ZoneOperation{Type: specs.Intersect, BodyID: specs.BodyID(100)},
				},
			}),
			BodyIDToShield:     map[specs.BodyID]ShieldBodyID{1: 1, 100: 2},
			MaterialIDToShield: map[specs.MaterialID]material.ShieldID{2: 200},
			Expected: []*zoneTree{
				&zoneTree{
					childrens:  []*zoneTree{},
					baseBodyID: 1,
					operations: []operation{operation{
						BodyID: 2,
						Type:   specs.Intersect,
					}},
					materialID: 200,
				},
			},
			ExpectedError: nil,
		})
	})

	t.Run("ManyZones", func(t *testing.T) {
		check(t, testCase{
			ZoneMap: createZoneMap(
				specs.Zone{
					ID:         specs.ZoneID(1),
					ParentID:   specs.ZoneID(0),
					BaseID:     specs.BodyID(1),
					MaterialID: specs.MaterialID(2),
					Construction: []specs.ZoneOperation{
						specs.ZoneOperation{Type: specs.Intersect, BodyID: specs.BodyID(100)},
					},
				},
				specs.Zone{
					ID:         specs.ZoneID(2),
					ParentID:   specs.ZoneID(1),
					BaseID:     specs.BodyID(300),
					MaterialID: specs.MaterialID(300),
				},
			),
			BodyIDToShield:     map[specs.BodyID]ShieldBodyID{1: 1, 100: 2, 300: 3},
			MaterialIDToShield: map[specs.MaterialID]material.ShieldID{2: 200, 300: 1},
			Expected: []*zoneTree{
				&zoneTree{
					childrens: []*zoneTree{
						&zoneTree{
							childrens:  []*zoneTree{},
							baseBodyID: 3,
							operations: []operation{},
							materialID: 1,
						},
					},
					baseBodyID: 1,
					operations: []operation{operation{
						BodyID: 2,
						Type:   specs.Intersect,
					}},
					materialID: 200,
				},
			},
			ExpectedError: nil,
		})
	})
}

// TODO remove all function with similiar signatures.
func createZoneMap(zones ...specs.Zone) map[specs.ZoneID]specs.Zone {
	res := map[specs.ZoneID]specs.Zone{}
	for _, z := range zones {
		res[z.ID] = z
	}
	return res
}
