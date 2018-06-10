package serialize

import (
	"fmt"

	"github.com/yaptide/yaptide/pkg/converter/setup"
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

func DetectorMarshaler(d setup.Detector) marshaler {
	return StructMarshaler(func(m fieldMarshaler) {
		m("id", Int64Marshaler(d.ID))
		m("geometry", func() (interface{}, error) {
			switch d := d.Geometry.(type) {
			case setup.DetectorGeomap:
				return DetectorGeomapMarshaler(d)()
			case setup.DetectorZones:
				return DetectorZonesMarshaler(d)()
			case setup.DetectorCylinder:
				return DetectorCylinderMarshaler(d)()
			case setup.DetectorMesh:
				return DetectorMeshMarshaler(d)()
			case setup.DetectorPlane:
				return DetectorPlaneMarshaler(d)()
			default:
				return nil, fmt.Errorf("unknown detector geometry")
			}
		})
		m("particle", ParticleMarshaler(d.ScoredParticle))
		m("scoring", DetectorScoringMarshaler(d.Scoring))
	})
}

func DetectorUnmarshaler(d *setup.Detector) unmarshaler {
	return StructUnmarshaler(func(u fieldUnmarshaler) {
		u("id", Int64Unmarshaler(&d.ID))
		u("geometry", UnionTypeUnmarshaler(
			func(unionType string) unmarshaler {
				switch unionType {
				case detectorType.geomap:
					return DetectorGeomapUnmarshaler(&d.Geometry)
				case detectorType.zone:
					return DetectorZonesUnmarshaler(&d.Geometry)
				case detectorType.cylinder:
					return DetectorCylinderUnmarshaler(&d.Geometry)
				case detectorType.mesh:
					return DetectorMeshUnmarshaler(&d.Geometry)
				case detectorType.plane:
					return DetectorPlaneUnmarshaler(&d.Geometry)
				default:
					return func(raw interface{}) error {
						return fmt.Errorf("unknown detector type")
					}
				}
			},
		))
		u("particle", ParticleUnmarshaler(&d.ScoredParticle))
		u("scoring", DetectorScoringUnmarshaler(&d.Scoring))
	})
}

func DetectorGeomapMarshaler(d setup.DetectorGeomap) marshaler {
	return StructMarshaler(func(m fieldMarshaler) {
		m("center", PointMarshaler(d.Center))
		m("size", Vec3DMarshaler(d.Size))
		m("slices", Vec3DIntMarshaler(d.Slices))
	})
}

func DetectorGeomapUnmarshaler(geometry *setup.DetectorGeometry) unmarshaler {
	return StructUnmarshaler(func(u fieldUnmarshaler) {
		d := setup.DetectorGeomap{}
		u("center", PointUnmarshaler(&d.Center))
		u("size", Vec3DUnmarshaler(&d.Size))
		u("slices", Vec3DIntUnmarshaler(&d.Slices))
		*geometry = d
	})
}

func DetectorZonesMarshaler(d setup.DetectorZones) marshaler {
	return StructMarshaler(func(m fieldMarshaler) {
		m("zones", ListMarshaler(d.Zones, Int64Marshaler))
	})
}

func DetectorZonesUnmarshaler(geometry *setup.DetectorGeometry) unmarshaler {
	return StructUnmarshaler(func(u fieldUnmarshaler) {
		d := setup.DetectorZones{}
		u("zones", ListUnmarshaler(&d.Zones, Int64Unmarshaler))
		*geometry = d
	})
}

func DetectorCylinderMarshaler(d setup.DetectorCylinder) marshaler {
	return StructMarshaler(func(m fieldMarshaler) {
		m("radius", RangeMarshaler(d.Radius))
		m("angle", RangeMarshaler(d.Angle))
		m("zValue", RangeMarshaler(d.Angle))
		m("slices", Vec3DCylindricalIntMarshaler(d.Slices))
	})
}

func DetectorCylinderUnmarshaler(geometry *setup.DetectorGeometry) unmarshaler {
	return StructUnmarshaler(func(u fieldUnmarshaler) {
		d := setup.DetectorCylinder{}
		u("radius", RangeUnmarshaler(&d.Radius))
		u("angle", RangeUnmarshaler(&d.Angle))
		u("zValue", RangeUnmarshaler(&d.Angle))
		u("slices", Vec3DCylindricalIntUnmarshaler(&d.Slices))
		*geometry = d
	})
}

func DetectorMeshMarshaler(d setup.DetectorMesh) marshaler {
	return StructMarshaler(func(m fieldMarshaler) {
		m("center", PointMarshaler(d.Center))
		m("size", Vec3DMarshaler(d.Size))
		m("slices", Vec3DIntMarshaler(d.Slices))
	})
}

func DetectorMeshUnmarshaler(geometry *setup.DetectorGeometry) unmarshaler {
	return StructUnmarshaler(func(u fieldUnmarshaler) {
		d := setup.DetectorMesh{}
		u("center", PointUnmarshaler(&d.Center))
		u("size", Vec3DUnmarshaler(&d.Size))
		u("slices", Vec3DIntUnmarshaler(&d.Slices))
		*geometry = d
	})
}

func DetectorPlaneMarshaler(d setup.DetectorPlane) marshaler {
	return StructMarshaler(func(m fieldMarshaler) {
		m("point", PointMarshaler(d.Point))
		m("normal", Vec3DMarshaler(d.Normal))
	})
}

func DetectorPlaneUnmarshaler(geometry *setup.DetectorGeometry) unmarshaler {
	return StructUnmarshaler(func(u fieldUnmarshaler) {
		d := setup.DetectorPlane{}
		u("point", PointUnmarshaler(&d.Point))
		u("normal", Vec3DUnmarshaler(&d.Normal))
		*geometry = d
	})
}

func DetectorScoringMarshaler(d setup.DetectorScoring) marshaler {
	return func() (interface{}, error) {
		switch d := d.(type) {
		case setup.PredefinedScoring:
			return StructMarshaler(func(m fieldMarshaler) {
				m("type", StringMarshaler(string(d)))
			})()
		case setup.LetTypeScoring:
			return StructMarshaler(func(m fieldMarshaler) {
				m("type", StringMarshaler(d.Type))
				m("material", Int64Marshaler(d.Material))
			})()
		default:
			return nil, fmt.Errorf("unknown detector scoring type %v", d)
		}
	}
}

func DetectorScoringUnmarshaler(d *setup.DetectorScoring) unmarshaler {
	return UnionTypeUnmarshaler(func(unionType string) unmarshaler {
		return func(raw interface{}) error {
			if predefinedScoringTypes[unionType] {
				*d = setup.PredefinedScoring(unionType)
			} else if letScoringTypes[unionType] {
				return StructUnmarshaler(func(u fieldUnmarshaler) {
					scoring := setup.LetTypeScoring{}
					u("type", StringUnmarshaler(&scoring.Type))
					u("material", Int64Unmarshaler(&scoring.Material))
					*d = scoring
				})(raw)
			} else {
				return fmt.Errorf("invalid scoring type")
			}
			return nil
		}
	})
}
