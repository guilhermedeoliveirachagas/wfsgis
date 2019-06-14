package ogc

type Bbox struct {

	Crs string `json:"CRS,omitempty"`
	Coords []float64 `json:"bbox,omitempty"`

}

func NewBbox(west float64,north float64,east float64,south float64) (*Bbox) {

	return &Bbox{Coords: []float64{west,north,east,south}, Crs: "http://www.opengis.net/def/crs/OGC/1.3/CRS84"}

}