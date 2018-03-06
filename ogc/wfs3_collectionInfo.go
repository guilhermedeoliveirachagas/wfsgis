package ogc

type CollectionInfo struct{

	Name string `json:"name"`
	Links []*Link `json:"links"`
	Title string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Extent *Bbox `json:"extent,omitempty"`
	CRS []string `json:"crs,omitempty"`
}


//initialize the slices
func NewCollectionInfo() *CollectionInfo{
	return &CollectionInfo{Links: []*Link{}, CRS: []string{"http://www.opengis.net/def/crs/OGC/1.3/CRS84"}}
}