package ogc

import (
	"encoding/json"

	"github.com/paulmach/orb/geojson"
	"github.com/paulmach/orb"
)

// type When struct {
// 	Type     string     `json:"@type,omitempty"`
// 	Datetime *time.Time `json:"datetime,omitempty"`
// }

type Feature struct {
	ID         string             `json:"id,omitempty"`
	Type       string             `json:"type,omitempty"`
	Geometry   orb.Geometry       `json:"geometry,omitempty"`
	Properties geojson.Properties `json:"properties,omitempty"`
	// When       *When              `json:"when,omitempty"`
}

func (f *Feature) UnmarshalJSON(b []byte) error {
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
	// if fg.When != nil {
	// 	f.When = &When{}
	// 	f.When.Type = fg.When.Type
	// 	f.When.Datetime = fg.When.Datetime
	// }

	return nil
}

func (f *Feature) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	m["id"] = f.ID
	m["type"] = f.Type
	m["geometry"] = geojson.NewGeometry(f.Geometry)
	m["properties"] = f.Properties
	// if f.When != nil {
	// 	m["when"] = f.When
	// }
	return json.Marshal(m)
}

func NewFeature(geometry orb.Geometry) *Feature {
	return &Feature{
		Type:       "Feature",
		Geometry:   geometry,
		Properties: make(map[string]interface{}),
	}
}
