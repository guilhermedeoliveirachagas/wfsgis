package ogc

// A FeatureCollection correlates to a GeoJSON feature collection.
type FeatureCollection struct {
	Type     string     `json:"type"`
	Features []*Feature `json:"features"`
}

func NewFeatureCollection() *FeatureCollection {
	return &FeatureCollection{
		Type:     "FeatureCollection",
		Features: []*Feature{},
	}
}
