package ogc

type GetFeatureRequest struct {
	Extent         *Bbox  `json:"extent,omitempty"`
	FeatureId      string `json:"featureId,omitempty"`
	CollectionName string `json:"collectionName,omitempty"`
}
