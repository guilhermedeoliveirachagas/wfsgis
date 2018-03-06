package handlers

import (
	"net/http"
	"github.com/boundlessgeo/wt/model"
	"github.com/gin-gonic/gin"
)

type ContentHandler struct {

	//stuff...

}

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
		cis := db.AllCollectionInfos()
		c.JSON(http.StatusOK, gin.H{"collections": cis})
	}
}
