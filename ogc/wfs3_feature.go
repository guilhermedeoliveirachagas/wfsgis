package ogc

import (
	"time"
	"github.com/paulmach/orb"
	"github.com/flaviostutz/orb/geojson"
)

type When struct {
	Type         string           `json:"@type,omitempty"`
	Datetime     *time.Time       `json:"datetime,omitempty"`
}

type Feature struct {
	ID         string             `json:"id,omitempty"`
	Type       string             `json:"type,omitempty"`
	Geometry   orb.Geometry       `json:"geometry,omitempty"`
	Properties geojson.Properties `json:"properties,omitempty"`
	When       *When              `json:"when,omitempty"`
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
	if fg.When != nil {
		f.When = &When{}
		f.When.Type = fg.When.Type
		f.When.Datetime = fg.When.Datetime
	}

	return nil
}

// func (f *Feature) MarshalJSON() ([]byte, error) {
// 	var b []byte
// 	b = append(b, `{"type":"`...)
// 	b = append(b, f.Type...)
// 	b = append(b, `",`...)
// 	b = append(b, `"id":"`...)
// 	b = append(b, f.ID...)
// 	b = append(b, `","properties":`...)
// 	p, _ := json.Marshal(f.Properties)
// 	b = append(b, p...)

// 	b = append(b, `,"geometry":`...)
// 	g, err := geojson.NewGeometry(f.Geometry).MarshalJSON()
// 	if err != nil {
// 		return nil, err
// 	}
// 	b = append(b, g...)
// 	b = append(b, `}`...)

// 	if f.When != nil {
// 		b = append(b, `,"when":`...)
// 		w, err := json.Marshal(f.When)
// 		if err != nil {
// 			return nil, err
// 		}
// 		b = append(b, w...)
// 		b = append(b, `}`...)
// 	}

// 	return b, nil
// }

func NewFeature(geometry orb.Geometry) *Feature {
	return &Feature{
		Type:       "Feature",
		Geometry:   geometry,
		Properties: make(map[string]interface{}),
	}
}
