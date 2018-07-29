package geometry

import (
	"github.com/yaptide/yaptide/pkg/converter/shield/material"
	"github.com/yaptide/yaptide/pkg/converter/specs"
)

// Geometry represent ready to serialize data for geo.dat file.
type Geometry struct {
	Bodies              []Body
	Zones               []Zone
	ZoneToMaterialPairs []ZoneToMaterial
}

// ConvertGeometry ...
func ConvertGeometry(
	bodyMap map[specs.BodyID]specs.Body,
	zoneMap map[specs.ZoneID]specs.Zone,
	materialIDToShield map[specs.MaterialID]material.ShieldID,
) (Geometry, map[specs.BodyID]ShieldBodyID, error) {
	bodies, bodyIDToShield, err := convertBodies(bodyMap)
	if err != nil {
		return Geometry{}, bodyIDToShield, err
	}

	bodiesWithBlackhole, blackholeBodyID, err := appendBlackholeBody(bodies)
	if err != nil {
		return Geometry{}, bodyIDToShield, err
	}

	zoneForest, err := convertZonesToZoneTreeForest(zoneMap, materialIDToShield, bodyIDToShield)
	if err != nil {
		return Geometry{}, bodyIDToShield, err
	}

	root := surroundZoneForestWithBlackholeZone(zoneForest, blackholeBodyID)

	zones, zoneToMaterialPairs, err := convertTreeToZones(root)
	if err != nil {
		return Geometry{}, bodyIDToShield, err
	}

	return Geometry{
		Bodies:              bodiesWithBlackhole,
		Zones:               zones,
		ZoneToMaterialPairs: zoneToMaterialPairs,
	}, bodyIDToShield, nil

}
