package ogc

import (
	"github.com/paulmach/orb/geojson"
	"github.com/paulmach/orb"
)

type Feature struct{

	ID string `json:"id"`
	Type string `json:"type"`
	Geometry orb.Geometry `json:"geometry"`
	Properties geojson.Properties `json:"properties"`
	Links []*Link `json:"links,omitempty"`
}

func NewFeature(geometry orb.Geometry) *Feature {
	return &Feature{
		Type:       "Feature",
		Geometry:   geometry,
		Properties: make(map[string]interface{}),
	}
}