package ogc

type Exception struct {
	Code        string `json:"code,omitempty"`
	Description string `json:"description,omitempty"`
}
