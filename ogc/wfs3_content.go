package ogc

type Content struct {
	Links       []*Link           `json:"links,omitempty"`
	Collections []*CollectionInfo `json:"Collections,omitempty"`
}

func NewContent() (*Content) {
	return &Content{Links: []*Link{}, Collections: []*CollectionInfo{}}
}

func (c *Content) AddCollection(info *CollectionInfo) {

	c.Collections = append(c.Collections, info)
}

func (c *Content) AddLink(link *Link) {

	c.Links = append(c.Links, link)
}
