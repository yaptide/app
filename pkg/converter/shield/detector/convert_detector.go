package detector

import (
	"bytes"
	"fmt"
	"sort"
	"unicode"

	"github.com/yaptide/yaptide/pkg/converter/geometry"
	"github.com/yaptide/yaptide/pkg/converter/shield/mapping"
	"github.com/yaptide/yaptide/pkg/converter/shield/material"
	"github.com/yaptide/yaptide/pkg/converter/specs"
)

// Detector represent specs.Detector,
type Detector struct {
	ScoringType string

	// Argument can be int64, float64 or string.
	Arguments []interface{}
}

// ConvertDetectors ...
func ConvertDetectors(
	detectorsMap map[specs.DetectorID]specs.Detector,
	materialIDToShield map[specs.MaterialID]material.ShieldID,
) ([]Detector, map[string]specs.DetectorID, error) {
	result := []Detector{}
	detectIds := []specs.DetectorID{}
	for k := range detectorsMap {
		detectIds = append(detectIds, k)
	}
	sort.SliceStable(detectIds, func(i, j int) bool { return detectIds[i] < detectIds[j] })

	detectorConverter := detectorConverter{materialIDToShield}

	uniqNameSet := map[string]specs.DetectorID{}
	mapFilenameToDetectorID := map[string]specs.DetectorID{}
	for n, id := range detectIds {
		specsDetector := detectorsMap[id]

		duplicateID, foundDuplicate := uniqNameSet[string(specsDetector.ID)]
		if foundDuplicate {
			return nil, uniqNameSet,
				fmt.Errorf(
					"Found name duplicates: \"%v\" for detector Ids: %d and %d",
					specsDetector.ID, id, duplicateID,
				)
		}
		uniqNameSet[string(specsDetector.ID)] = specsDetector.ID

		filename := createDetectorFileName(string(specsDetector.ID), n)
		mapFilenameToDetectorID[filename] = specsDetector.ID

		detector, err := detectorConverter.convertDetector(&specsDetector, filename)
		if err != nil {
			return nil, nil, err
		}

		result = append(result, detector)
	}
	return result, mapFilenameToDetectorID, nil
}

type detectorConverter struct {
	materialIDToShield map[specs.MaterialID]material.ShieldID
}

func (d detectorConverter) convertDetector(
	detect *specs.Detector, filename string,
) (Detector, error) {
	switch geo := detect.Geometry.(type) {
	case specs.DetectorGeomap:
		return Detector{}, fmt.Errorf("Geomap detector serialization not implemented")
	case specs.DetectorZones:
		return Detector{}, fmt.Errorf("Zone detector serialization not implemented")

	case specs.DetectorCylinder:
		return d.convertStandardGeometryDetector(detect, filename)
	case specs.DetectorMesh:
		return d.convertStandardGeometryDetector(detect, filename)
	case specs.DetectorPlane:
		return d.convertStandardGeometryDetector(detect, filename)

	default:
		return Detector{}, fmt.Errorf("Unknown detector type: %T", geo)
	}
}

func (d detectorConverter) convertStandardGeometryDetector(
	detect *specs.Detector, filename string,
) (Detector, error) {
	var newDetector Detector

	switch geo := detect.Geometry.(type) {
	case specs.DetectorCylinder:
		newDetector = Detector{
			ScoringType: "CYL",
			Arguments: []interface{}{
				geo.Radius.Min,
				geo.Angle.Min,
				geo.ZValue.Min,
				geo.Radius.Max,
				geo.Angle.Max,
				geo.ZValue.Max,

				geo.Slices.Radius,
				geo.Slices.Angle,
				geo.Slices.Z,
			},
		}
	case specs.DetectorMesh:
		xMin, xMax := geometry.CenterAndSizeToMinAndMax(geo.Center.X, geo.Size.Y)
		yMin, yMax := geometry.CenterAndSizeToMinAndMax(geo.Center.Y, geo.Size.Y)
		zMin, zMax := geometry.CenterAndSizeToMinAndMax(geo.Center.Z, geo.Size.Z)
		newDetector = Detector{
			ScoringType: "MSH",
			Arguments: []interface{}{
				xMin,
				yMin,
				zMin,
				xMax,
				yMax,
				zMax,

				geo.Slices.X,
				geo.Slices.Y,
				geo.Slices.Z,
			},
		}
	case specs.DetectorPlane:
		newDetector = Detector{
			ScoringType: "PLANE",
			Arguments: []interface{}{
				geo.Point.X,
				geo.Point.Y,
				geo.Point.Z,
				geo.Normal.X,
				geo.Normal.Y,
				geo.Normal.Z,
				"",
				"",
				"",
			},
		}

	}

	particleInShieldFormat, err := mapping.ParticleToShield(detect.ScoredParticle)
	if err != nil {
		return Detector{}, fmt.Errorf("%s", err.Error())
	}

	scoringInShield, err := mapping.ScoringToShield(detect.Scoring)
	if err != nil {
		return Detector{}, fmt.Errorf("%s", err.Error())
	}

	newDetector.Arguments = append(newDetector.Arguments,
		particleInShieldFormat,
		scoringInShield,
		filename,
	)

	newDetector.Arguments, err = d.appendHeavyIonOrLetfluCard(
		newDetector.Arguments, detect.ScoredParticle, detect.Scoring,
	)
	if err != nil {
		return Detector{}, fmt.Errorf("%s", err.Error())
	}
	return newDetector, nil
}

// TODO: we need A and Z if partile is not HeavyIon and scoring is LetTypeScoring
func (d detectorConverter) appendHeavyIonOrLetfluCard(
	arguments []interface{}, particle specs.Particle, scoringType specs.DetectorScoring,
) ([]interface{}, error) {
	switch part := particle.(type) {
	case specs.HeavyIon:
		arguments = append(arguments, part.NucleonsCount, part.Charge)
		switch scoring := scoringType.(type) {
		case specs.LetTypeScoring:
			material, found := d.materialIDToShield[scoring.Material]
			if !found {
				return nil, fmt.Errorf("Can not found Material{ID: %d} for LetTypeScoring", scoring.Material)
			}
			arguments = append(arguments, int64(material))
		default:
			arguments = append(arguments, "")
		}
		return append(arguments, "", "", ""), nil
	default:
		return arguments, nil
	}
}

func createDetectorFileName(name string, detectorN int) string {
	buff := &bytes.Buffer{}

	for _, c := range name {
		switch {
		case unicode.IsDigit(c):
			buff.WriteRune(c)
		case c <= unicode.MaxASCII && unicode.IsLetter(c):
			buff.WriteRune(unicode.ToLower(c))
		default:
			buff.WriteString("_")
		}
	}

	fmt.Fprintf(buff, "%d", detectorN)
	return buff.String()
}
