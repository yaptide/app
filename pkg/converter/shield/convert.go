package shield

import (
	"fmt"

	"github.com/yaptide/yaptide/pkg/converter"
	"github.com/yaptide/yaptide/pkg/converter/shield/detector"
	"github.com/yaptide/yaptide/pkg/converter/shield/geometry"
	"github.com/yaptide/yaptide/pkg/converter/shield/material"
	"github.com/yaptide/yaptide/pkg/converter/specs"
)

// Convert simulation specs model to easily serializable data,
// which is input for shield serializer.
// Return error, if specs data are not semantically correct.
func Convert(simulationSpecs converter.Specs) (RawShieldSpecs, SerializationContext, error) {
	simContext := NewSerializationContext()

	err := checkSpecsCompleteness(simulationSpecs)
	if err != nil {
		return RawShieldSpecs{}, simContext, err
	}

	mapMaterials := map[specs.MaterialID]specs.Material{}
	for _, material := range simulationSpecs.Materials {
		mapMaterials[material.ID] = material
	}

	mapBodies := map[specs.BodyID]specs.Body{}
	for _, body := range simulationSpecs.Bodies {
		mapBodies[body.ID] = body
	}

	mapZones := map[specs.ZoneID]specs.Zone{}
	for _, zone := range simulationSpecs.Zones {
		mapZones[zone.ID] = zone
	}

	mapDetectors := map[specs.DetectorID]specs.Detector{}
	for _, detector := range simulationSpecs.Detectors {
		mapDetectors[detector.ID] = detector
	}

	materials, materialIDToShield, materialErr := material.ConvertMaterials(mapMaterials)
	geometry, mapBodyToShield, geometryErr := geometry.ConvertGeometry(
		mapBodies, mapZones, materialIDToShield,
	)
	detectors, mapDetectorToShield, detectorErr := detector.ConvertDetectors(
		mapDetectors, materialIDToShield,
	)

	for key, value := range materialIDToShield {
		simContext.MapMaterialID[value] = key
	}

	for key, value := range mapBodyToShield {
		simContext.MapBodyID[value] = key
	}

	simContext.MapFilenameToDetectorID = mapDetectorToShield

	_, _, _ = materialErr, geometryErr, detectorErr

	return RawShieldSpecs{
			Materials: materials,
			Geometry:  geometry,
			Detectors: detectors,
			Beam:      simulationSpecs.Beam,
			Options:   simulationSpecs.Options,
		},
		simContext,
		nil
}

func checkSpecsCompleteness(specs converter.Specs) error {
	createMissingError := func(mapName string) error {
		return fmt.Errorf("[serializer]: %s map is null", mapName)
	}

	createEmptyError := func(mapName string) error {
		return fmt.Errorf("[serializer]: %s map is empty", mapName)
	}

	switch {
	case specs.Bodies == nil:
		return createMissingError("Bodies")
	case specs.Zones == nil:
		return createMissingError("Zones")
	case specs.Materials == nil:
		return createMissingError("Materials")
	case specs.Detectors == nil:
		return createMissingError("Detectors")
	}

	switch {
	case len(specs.Bodies) == 0:
		return createEmptyError("Bodies")
	case len(specs.Zones) == 0:
		return createEmptyError("Zones")
	case len(specs.Materials) == 0:
		return createEmptyError("Materials")
	case len(specs.Detectors) == 0:
		return createEmptyError("Detectors")
	}

	return nil
}
