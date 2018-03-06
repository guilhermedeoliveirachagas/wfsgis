package ogc

type Conformance struct {
	ConformsTo []string `json:"conformsTo"`
}

func NewConformance() *Conformance {
	return &Conformance{ConformsTo: []string{"http://www.opengis.net/spec/wfs-1/3.0/req/core", "http://www.opengis.net/spec/wfs-1/3.0/req/geojson", "http://www.opengis.net/spec/wfs-1/3.0/req/oas30"}}
}
