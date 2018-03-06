package handlers

import (
	"net/http"
	"github.com/boundlessgeo/wt/model"
	"github.com/boundlessgeo/wt/ogc"
	"github.com/gin-gonic/gin"
)

func (h *HTTPServer) makeContentHandlers(d *model.DB) {


	//handlers
	conformance := ConformanceHandler{}
	feature := FeatureHandler{Store: d}

	h.router.GET("/", getCollections(d))
	//the base endpoint should provide a list of all the supported collections
	// aka tables
	h.router.GET("/collection/:collectionId/", feature.Handle)
	//Conformance endpoint
	h.router.GET("/api/conformance", conformance.Handle)
	h.router.NoRoute(feature.Handle)
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
