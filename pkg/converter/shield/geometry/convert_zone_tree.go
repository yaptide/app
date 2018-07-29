package geometry

import (
	"fmt"

	"github.com/yaptide/yaptide/pkg/converter/shield/material"
	"github.com/yaptide/yaptide/pkg/converter/specs"
)

type operation struct {
	BodyID ShieldBodyID
	Type   specs.ZoneOperationType
}

type zoneTree struct {
	childrens []*zoneTree

	baseBodyID ShieldBodyID
	operations []operation

	materialID material.ShieldID
}

func convertZonesToZoneTreeForest(
	zoneMap map[specs.ZoneID]specs.Zone,
	materialIDToShield map[specs.MaterialID]material.ShieldID,
	bodyIDToShield map[specs.BodyID]ShieldBodyID) ([]*zoneTree, error) {

	converter := zoneConverter{
		zoneMap:            zoneMap,
		materialIDToShield: materialIDToShield,
		bodyIDToShield:     bodyIDToShield,
	}
	return converter.convertZonesToZoneTreeForest()
}

type zoneConverter struct {
	zoneMap            map[specs.ZoneID]specs.Zone
	materialIDToShield map[specs.MaterialID]material.ShieldID
	bodyIDToShield     map[specs.BodyID]ShieldBodyID
}

func (z *zoneConverter) convertZonesToZoneTreeForest() ([]*zoneTree, error) {
	forest := []*zoneTree{}

	for _, zoneModel := range z.zoneMap {
		if zoneModel.ParentID == specs.RootID {
			newZoneTree, err := z.createZoneTree(&zoneModel)
			if err != nil {
				return nil, err
			}
			forest = append(forest, newZoneTree)
		}
	}
	return forest, nil
}

func (z *zoneConverter) createZoneTree(zoneModel *specs.Zone) (*zoneTree, error) {
	baseBodyID, found := z.bodyIDToShield[zoneModel.BaseID]
	if !found {
		return nil, fmt.Errorf("Cannot find body: %d", zoneModel.BaseID)
	}

	operations, err := z.convertZoneOperations(zoneModel.Construction)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	materialID, found := z.materialIDToShield[zoneModel.MaterialID]
	if !found {
		return nil, fmt.Errorf("Cannot find material: %d", zoneModel.MaterialID)
	}

	childModelIDs := []specs.ZoneID{}
	for _, zone := range z.zoneMap {
		if zone.ParentID == zoneModel.ID {
			childModelIDs = append(childModelIDs, zone.ID)
		}
	}

	childrens := []*zoneTree{}
	for _, childModelID := range childModelIDs {
		childModel, found := z.zoneMap[childModelID]
		if !found {
			return nil, fmt.Errorf("Can not find Children {ID: %d}", childModelID)
		}

		child, err := z.createZoneTree(&childModel)
		if err != nil {
			return nil, err
		}

		childrens = append(childrens, child)
	}

	return &zoneTree{
		childrens:  childrens,
		baseBodyID: baseBodyID,
		operations: operations,
		materialID: materialID,
	}, nil
}

func (z *zoneConverter) convertZoneOperations(
	specsOperations []specs.ZoneOperation,
) ([]operation, error) {
	operations := []operation{}
	for _, o := range specsOperations {
		bodyID, found := z.bodyIDToShield[o.BodyID]
		if !found {
			return nil, fmt.Errorf("Cannot find body: %d", o.BodyID)
		}
		operations = append(operations, operation{
			BodyID: bodyID,
			Type:   o.Type,
		})
	}
	return operations, nil
}

func surroundZoneForestWithBlackholeZone(
	zoneForest []*zoneTree, blackholeBodyID ShieldBodyID,
) *zoneTree {
	return &zoneTree{
		childrens:  zoneForest,
		baseBodyID: blackholeBodyID,
		operations: []operation{},
	}
}
