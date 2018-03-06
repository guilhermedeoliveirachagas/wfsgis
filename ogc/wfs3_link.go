package ogc

import (

)

type Link struct{

	Href string `json:"href"`
	Rel string `json:"rel,omitempty"`
	Type string `json:"rel,omitempty"`
	Reflang string `json:"hreflang,omitempty"`
	Title string `json:"title,omitempty"`

}