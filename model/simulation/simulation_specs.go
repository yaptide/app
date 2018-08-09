package simulation

import (
	"encoding/json"

	"github.com/yaptide/yaptide/pkg/converter"
	"gopkg.in/mgo.v2/bson"
)

// SimulationSpecs ...
type SpecsDB struct {
	ID                bson.ObjectId `json:"id" bson:"_id"`
	UserID            bson.ObjectId `json:"userId" bson:"userId"`
	Specs             `json:"specs" bson:"specs"`
	serializeOverride `json:",-" bson:"-"`
}

type Specs struct {
	Materials []Material
	Bodies    []Body
	Zones     []Zone
	Detectors []Detector
	Beam      Beam
	Options   SimulationOptions
}

type Material struct {
	Name string
	converter.Material
}

type Body = converter.Body
type Zone struct {
	Name  string
	Color Color
	converter.Zone
}
type Detector struct {
	Name string
	converter.Detector
}
type Beam = converter.Beam
type SimulationOptions = converter.SimulationOptions

func (c Specs) ConverterConfiguration() converter.Specs {
	materials := make([]converter.Material, len(c.Materials))
	for i, m := range c.Materials {
		materials[i] = m.Material
	}
	detectors := make([]converter.Detector, len(c.Detectors))
	for i, d := range c.Detectors {
		detectors[i] = d.Detector
	}
	zones := make([]converter.Zone, len(c.Zones))
	for i, z := range c.Zones {
		zones[i] = z.Zone
	}

	return converter.Specs{
		Materials: materials,
		Bodies:    c.Bodies,
		Zones:     zones,
		Detectors: detectors,
		Beam:      c.Beam,
		Options:   c.Options,
	}
}

// GetBSON ...
func (c Specs) GetBSON() (interface{}, error) {
	return specsMarshaler(c)()
}

// SetBSON ...
func (c *Specs) SetBSON(raw bson.Raw) error {
	var unpack map[string]interface{}
	if err := raw.Unmarshal(&unpack); err != nil {
		return err
	}
	return specsUnmarshaler(c)(unpack)
}

func (c Specs) MarshalJSON() ([]byte, error) {
	raw, err := specsMarshaler(c)()
	if err != nil {
		return nil, err
	}
	return json.Marshal(raw)
}

func (c *Specs) UnmarshalJSON(raw []byte) error {
	var unpack interface{}
	if err := json.Unmarshal(raw, &unpack); err != nil {
		return err
	}
	return specsUnmarshaler(c)(unpack)
}

type serializeOverride interface {
	GetBSON() (interface{}, error)
	SetBSON(bson.Raw) error
	MarshalJSON() ([]byte, error)
	UnmarshalJSON([]byte) error
}
