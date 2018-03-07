package ogc

type GetFeatureRequest struct {
	Extent         *Bbox
	FeatureId      string
	CollectionName string
}
