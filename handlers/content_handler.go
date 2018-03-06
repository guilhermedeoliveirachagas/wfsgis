package handlers

import (
	"github.com/boundlessgeo/wt/model"
	"github.com/gin-gonic/gin"
)

type ContentHandler struct {

	//stuff...

}

func (h *HTTPServer) makeContentHandlers(d *model.DB) {
	h.router.GET("/", getCollections(d))
}

func getCollections(db *model.DB) func(*gin.Context) {
	return func(g *gin.Context) {

	}
}
