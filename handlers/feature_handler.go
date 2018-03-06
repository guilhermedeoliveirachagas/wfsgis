package handlers

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/boundlessgeo/wt/model"
	"github.com/boundlessgeo/wt/ogc"
	"github.com/gin-gonic/gin"
)


type FeatureHandler struct {
	Store *model.DB
}

func (h *HTTPServer) makeFeatureHandlers(d *model.DB) {

	h.router.GET("/collections/:collid/items", getFeatures(d))
	h.router.GET("/collections/:collid/items/:itemid", getFeatureFromCollection(d))
	h.router.POST("/collections/:collid/items", createFeature(d))
	h.router.PUT("/collections/:collid/items/:fid", updateFeature(d))
	h.router.DELETE("/collections/:collid/items/:fid", deleteFeature(d))

}

/**
Updates a feature
 */
func updateFeature(db *model.DB) func(*gin.Context){
	return nil
}
/**
Creates a feature
 */
func createFeature(db *model.DB) func(*gin.Context){
	return nil
}
/**
Deletes a feature
 */
func deleteFeature(db *model.DB) func(*gin.Context){
	return nil
}
/**
Gets a feature by collection
 */
func getFeatureFromCollection(db *model.DB) func(*gin.Context){
	return nil
}
/**
Gets features
 */
func getFeatures(db *model.DB) func(*gin.Context){
	return nil
}



func (fh *FeatureHandler) Handle(c *gin.Context) {

	switch c.Request.Method {
	case "GET":
		{
			collectionName := strings.TrimLeft(c.Request.URL.Path, "/")
			getFeature := ogc.GetFeatureRequest{Extent: ogc.NewBbox(-180, 90, 180, -90), FeatureId: "", CollectionName: collectionName}
			fh.Store.GetFeatures(getFeature)

		}
	case "POST":
		{
			collectionName := strings.TrimLeft(c.Request.URL.Path, "/")
			fc := &ogc.FeatureCollection{}
			data, _ := ioutil.ReadAll(c.Request.Body)
			json.Unmarshal(data, fc)
			fh.Store.CreateCollectionTable(collectionName, fc.Features)
		}

	default:
		{
			c.JSON(405, ogc.Exception{"405", "Method not allowed"})
		}
	}

}
