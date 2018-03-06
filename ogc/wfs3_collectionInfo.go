package ogc

type CollectionInfo struct {
	Name        string   `json:"name" db:"table_name"`
	Links       []*Link  `json:"links" db:"links"`
	Title       string   `json:"title,omitempty" db:"title"`
	Description string   `json:"description,omitempty" db:"description"`
	Extent      *Bbox    `json:"extent,omitempty" db:"extent"`
	CRS         []string `json:"crs,omitempty" db:"crs"`
}

//initialize the slices
func NewCollectionInfo() *CollectionInfo {
	return &CollectionInfo{Links: []*Link{}, CRS: []string{"http://www.opengis.net/def/crs/OGC/1.3/CRS84"}}
}
