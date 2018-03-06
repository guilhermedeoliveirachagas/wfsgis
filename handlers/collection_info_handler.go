package handlers

import (
	"net/http"

	"github.com/boundlessgeo/wt/model"
	"github.com/boundlessgeo/wt/ogc"
	"github.com/gin-gonic/gin"
)

func (h *HTTPServer) makeCollectionHandlers(d *model.DB) {

	h.router.GET("/collections", getCollectionsInfo(d))
	h.router.GET("/collections/:collid/schema", getCollectionInfo(d))
	h.router.PUT("/collections/:collid/schema", updateCollectionInfo(d))
	h.router.POST("/collections", createCollectionInfo(d))
	h.router.DELETE("/collections/:collid", deleteCollection(d))

}
/**
Deletes a collection
 */
func deleteCollection(db *model.DB) func(*gin.Context){

}

/**
Creates a collection
 */
func createCollectionInfo(db *model.DB) func(*gin.Context){

}
/**
Updates a collection
 */
func updateCollectionInfo(db *model.DB) func(*gin.Context){

}

/**
Gets a single collection
 */
func getCollectionInfo(db *model.DB) func(*gin.Context) {

}

/**
Gets all collections
 */
func getCollectionsInfo(db *model.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		cidbs := db.AllCollectionInfos()
		cis := make([]*ogc.CollectionInfo, 0)
		for _, v := range cidbs {
			cis = append(cis, v.CollectionInfo)
		}
		c.JSON(http.StatusOK, gin.H{"collections": cis})
	}
}

func createCollection(db *model.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		db.CreateCollectionTable()
	}
}
