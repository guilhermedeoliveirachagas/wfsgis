package ogc

import (
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
)

type Feature struct {
	ID         string             `json:"id"`
	Type       string             `json:"type"`
	Geometry   orb.Geometry       `json:"geometry"`
	Properties geojson.Properties `json:"properties"`
}

func (f *Feature) UnmarshalJSON(b []byte) error {

	var fg *geojson.Feature
	fg, err := geojson.UnmarshalFeature(b)
	if err != nil {
		return err
	}

	f.Geometry = fg.Geometry
	if fg.ID == nil {
		f.ID = ""
	} else {
		f.ID = fg.ID.(string)
	}
	f.Properties = fg.Properties
	f.Type = fg.Type

	return nil
}

func NewFeature(geometry orb.Geometry) *Feature {
	return &Feature{
		Type:       "Feature",
		Geometry:   geometry,
		Properties: make(map[string]interface{}),
	}
}
