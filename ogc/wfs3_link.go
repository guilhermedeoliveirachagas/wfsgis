package ogc

type Link struct {
	Href    string `json:"href,omitempty"`
	Rel     string `json:"rel,omitempty"`
	Typ     string `json:"type,omitempty"`
	Reflang string `json:"hreflang,omitempty"`
	Title   string `json:"title,omitempty"`
}
