package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/boundlessgeo/wt/model"
	"github.com/boundlessgeo/wt/ogc"
	"github.com/gin-gonic/gin"
)

type FeatureHandler struct {
	Store *model.DB
}

func (h *HTTPServer) makeFeatureHandlers(d *model.DB) {

	h.router.GET("/collections/:collid/items", getFeatures(d))
	h.router.GET("/collections/:collid/items/:itemid", getFeatureById(d))
	h.router.POST("/collections/:collid/items", createFeature(d))
	h.router.PUT("/collections/:collid/items/:fid", updateFeature(d))
	h.router.DELETE("/collections/:collid/items/:fid", deleteFeature(d))

}

/**
Updates a feature
*/
func updateFeature(db *model.DB) func(*gin.Context) {
	return func(c *gin.Context) {

	}
}

/**
Creates a feature
*/
func createFeature(db *model.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		collectionName := c.Param("collid")

		var fc ogc.FeatureCollection
		data, _ := ioutil.ReadAll(c.Request.Body)
		err := json.Unmarshal(data, &fc)

		//if inputErr := c.BindJSON(&fc); inputErr != nil {
		//	c.JSON(http.StatusBadRequest, inputErr.Error)
		//	return
		//}

		_, err = db.InsertFeature(collectionName, fc.Features)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ogc.Exception{Code: "500", Description: "Error inserting feature"})
		}
		c.JSON(http.StatusCreated, fc)
	}
}

/**
Deletes a feature
*/
func deleteFeature(db *model.DB) func(*gin.Context) {
	return nil
}

/**
Gets a feature by id
*/
func getFeatureById(db *model.DB) func(*gin.Context) {
	return nil
}

/**
Gets features
*/
func getFeatures(db *model.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		//collectionName := c.Param("collid")
		//getFeature := ogc.GetFeatureRequest{Extent: ogc.NewBbox(-180, 90, 180, -90), FeatureId: "", CollectionName: collectionName}
		//features, err := db.GetFeatures(getFeature)
		//if err != nil{
		//
		//	c.JSON(500, ogc.Exception{"500","Error fetching features"})
		//}

		fc := ogc.NewFeatureCollection()
		//fc.Features = features
		fc.Type = "FeatureCollection"
		c.JSON(200, fc)

	}
}
