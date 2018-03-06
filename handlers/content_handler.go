package handlers

import (
	"net/http"

	"github.com/boundlessgeo/wt/model"
	"github.com/boundlessgeo/wt/ogc"
	"github.com/gin-gonic/gin"
)

func (h *HTTPServer) makeContentHandlers(d *model.DB) {
	h.router.GET("/", getCollections(d))
}

func getCollections(db *model.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		cidbs := db.AllCollectionInfos()
		cis := make([]*ogc.CollectionInfo, 0)
		for _, v := range cidbs {
			cis = append(cis, v.CollectionInfo)
		}
		c.JSON(http.StatusOK, gin.H{"collections": cis})
	}
}
