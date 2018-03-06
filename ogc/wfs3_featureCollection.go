package ogc

// A FeatureCollection correlates to a GeoJSON feature collection.
type FeatureCollection struct {
	Type     string     `json:"type"`
	Features []*Feature `json:"features"`
	Links    []*Link    `json:"links,omitifempty"`
}

func NewFeatureCollection() *FeatureCollection {
	return &FeatureCollection{
		Type:     "FeatureCollection",
		Features: []*Feature{},
	}
}
