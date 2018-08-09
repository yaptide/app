package simulation

import (
	"fmt"

	"github.com/yaptide/yaptide/pkg/converter/specs"
	. "github.com/yaptide/yaptide/pkg/serialize"
)

var detectorType = struct {
	geomap   string
	zone     string
	cylinder string
	mesh     string
	plane    string
}{
	geomap:   "geomap",
	zone:     "zone",
	cylinder: "cylinder",
	mesh:     "mesh",
	plane:    "plane",
}

var predefinedScoringTypes = map[string]bool{
	"dose":       true,
	"energy":     true,
	"fluence":    true,
	"avg_energy": true,
	"avg_beta":   true,
	"spc":        true,
	"alanine":    true,
	"counter":    true,
	"ddd":        true,
	"crossflu":   true,
}

// TODO: write test checking if all values are assigned
// TODO: not sure about that
var letScoringTypes = map[string]bool{
	"letflu": true,
	"dlet":   true,
	"tlet":   true,
}

func detectorMarshaler(d Detector) marshaler {
	return StructMarshaler(func(m fieldMarshaler) {
		m("id", Int64Marshaler(int64(d.ID)))
		m("name", StringMarshaler(d.Name))
		m("geometry", func() (interface{}, error) {
			switch d := d.Geometry.(type) {
			case specs.DetectorGeomap:
				return detectorGeomapMarshaler(d)()
			case specs.DetectorZones:
				return detectorZonesMarshaler(d)()
			case specs.DetectorCylinder:
				return detectorCylinderMarshaler(d)()
			case specs.DetectorMesh:
				return detectorMeshMarshaler(d)()
			case specs.DetectorPlane:
				return detectorPlaneMarshaler(d)()
			case nil:
				return nil, nil
			default:
				return nil, fmt.Errorf("unknown detector geometry")
			}
		})
		m("particle", particleMarshaler(d.ScoredParticle))
		m("scoring", detectorScoringMarshaler(d.Scoring))
	})
}

func detectorUnmarshaler(d *Detector) unmarshaler {
	return StructUnmarshaler(func(u fieldUnmarshaler) {
		u("id", Int64Unmarshaler((*int64)(&d.ID)))
		u("name", StringUnmarshaler(&d.Name))
		u("geometry", UnionTypeUnmarshaler(
			func(unionType string) unmarshaler {
				switch unionType {
				case detectorType.geomap:
					return detectorGeomapUnmarshaler(&d.Geometry)
				case detectorType.zone:
					return detectorZonesUnmarshaler(&d.Geometry)
				case detectorType.cylinder:
					return detectorCylinderUnmarshaler(&d.Geometry)
				case detectorType.mesh:
					return detectorMeshUnmarshaler(&d.Geometry)
				case detectorType.plane:
					return detectorPlaneUnmarshaler(&d.Geometry)
				default:
					return func(raw interface{}) error {
						return fmt.Errorf("unknown detector type")
					}
				}
			},
		))
		u("particle", particleUnmarshaler(&d.ScoredParticle))
		u("scoring", detectorScoringUnmarshaler(&d.Scoring))
	})
}

func detectorGeomapMarshaler(d specs.DetectorGeomap) marshaler {
	return StructMarshaler(func(m fieldMarshaler) {
		m("type", StringMarshaler(detectorType.geomap))
		m("center", pointMarshaler(d.Center))
		m("size", vec3DMarshaler(d.Size))
		m("slices", vec3DIntMarshaler(d.Slices))
	})
}

func detectorGeomapUnmarshaler(geometry *specs.DetectorGeometry) unmarshaler {
	return StructUnmarshaler(func(u fieldUnmarshaler) {
		d := specs.DetectorGeomap{}
		u("center", pointUnmarshaler(&d.Center))
		u("size", vec3DUnmarshaler(&d.Size))
		u("slices", vec3DIntUnmarshaler(&d.Slices))
		*geometry = d
	})
}

func detectorZonesMarshaler(d specs.DetectorZones) marshaler {
	return StructMarshaler(func(m fieldMarshaler) {
		m("type", StringMarshaler(detectorType.zone))
		m("zones", ListMarshaler(d.Zones, Int64Marshaler))
	})
}

func detectorZonesUnmarshaler(geometry *specs.DetectorGeometry) unmarshaler {
	return StructUnmarshaler(func(u fieldUnmarshaler) {
		d := specs.DetectorZones{}
		u("zones", ListUnmarshaler(&d.Zones, Int64Unmarshaler))
		*geometry = d
	})
}

func detectorCylinderMarshaler(d specs.DetectorCylinder) marshaler {
	return StructMarshaler(func(m fieldMarshaler) {
		m("type", StringMarshaler(detectorType.cylinder))
		m("radius", rangeMarshaler(d.Radius))
		m("angle", rangeMarshaler(d.Angle))
		m("zValue", rangeMarshaler(d.ZValue))
		m("slices", vec3DCylindricalIntMarshaler(d.Slices))
	})
}

func detectorCylinderUnmarshaler(geometry *specs.DetectorGeometry) unmarshaler {
	return StructUnmarshaler(func(u fieldUnmarshaler) {
		d := specs.DetectorCylinder{}
		u("radius", rangeUnmarshaler(&d.Radius))
		u("angle", rangeUnmarshaler(&d.Angle))
		u("zValue", rangeUnmarshaler(&d.ZValue))
		u("slices", vec3DCylindricalIntUnmarshaler(&d.Slices))
		*geometry = d
	})
}

func detectorMeshMarshaler(d specs.DetectorMesh) marshaler {
	return StructMarshaler(func(m fieldMarshaler) {
		m("type", StringMarshaler(detectorType.mesh))
		m("center", pointMarshaler(d.Center))
		m("size", vec3DMarshaler(d.Size))
		m("slices", vec3DIntMarshaler(d.Slices))
	})
}

func detectorMeshUnmarshaler(geometry *specs.DetectorGeometry) unmarshaler {
	return StructUnmarshaler(func(u fieldUnmarshaler) {
		d := specs.DetectorMesh{}
		u("center", pointUnmarshaler(&d.Center))
		u("size", vec3DUnmarshaler(&d.Size))
		u("slices", vec3DIntUnmarshaler(&d.Slices))
		*geometry = d
	})
}

func detectorPlaneMarshaler(d specs.DetectorPlane) marshaler {
	return StructMarshaler(func(m fieldMarshaler) {
		m("type", StringMarshaler(detectorType.plane))
		m("point", pointMarshaler(d.Point))
		m("normal", vec3DMarshaler(d.Normal))
	})
}

func detectorPlaneUnmarshaler(geometry *specs.DetectorGeometry) unmarshaler {
	return StructUnmarshaler(func(u fieldUnmarshaler) {
		d := specs.DetectorPlane{}
		u("point", pointUnmarshaler(&d.Point))
		u("normal", vec3DUnmarshaler(&d.Normal))
		*geometry = d
	})
}

func detectorScoringMarshaler(d specs.DetectorScoring) marshaler {
	return func() (interface{}, error) {
		switch d := d.(type) {
		case specs.PredefinedScoring:
			return StructMarshaler(func(m fieldMarshaler) {
				m("type", StringMarshaler(string(d)))
			})()
		case specs.LetTypeScoring:
			return StructMarshaler(func(m fieldMarshaler) {
				m("type", StringMarshaler(d.Type))
				m("material", Int64Marshaler(int64(d.Material)))
			})()
		case nil:
			return nil, nil
		default:
			return nil, fmt.Errorf("unknown detector scoring type %v", d)
		}
	}
}

func detectorScoringUnmarshaler(d *specs.DetectorScoring) unmarshaler {
	return UnionTypeUnmarshaler(func(unionType string) unmarshaler {
		return func(raw interface{}) error {
			if predefinedScoringTypes[unionType] {
				*d = specs.PredefinedScoring(unionType)
			} else if letScoringTypes[unionType] {
				return StructUnmarshaler(func(u fieldUnmarshaler) {
					scoring := specs.LetTypeScoring{}
					u("type", StringUnmarshaler(&scoring.Type))
					u("material", Int64Unmarshaler((*int64)(&scoring.Material)))
					*d = scoring
				})(raw)
			} else {
				return fmt.Errorf("invalid scoring type")
			}
			return nil
		}
	})
}
