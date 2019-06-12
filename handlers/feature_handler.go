package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/boundlessgeo/wfs3/model"
	"github.com/boundlessgeo/wfs3/ogc"
	"github.com/gin-gonic/gin"
)

type FeatureHandler struct {
	Store *model.DB
}

func (h *HTTPServer) makeFeatureHandlers(d *model.DB) {

	h.router.GET("/collections/:collid/items", getFeatures(d))
	h.router.GET("/collections/:collid/items/:itemid", getFeatureById(d))
	h.router.POST("/collections/:collid/items", createFeature(d))
	h.router.PUT("/collections/:collid/items/:itemid", updateFeature(d))
	h.router.DELETE("/collections/:collid/items/:itemid", deleteFeature(d))

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
		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		log.Printf("createFeature %v %v", collectionName, fc.Features)
		if len(fc.Features) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"message": "No features found in json body"})
			return
		}
		_, err = db.InsertFeature(collectionName, fc.Features)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, fc)
	}
}

/**
Deletes a feature
*/
func deleteFeature(db *model.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		collid := c.Param("collid")
		itemid := c.Param("itemid")
		if err := db.DeleteItem(collid, itemid); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"success": "true"})
		}
	}
}

/**
Gets a feature by id
*/
func getFeatureById(db *model.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		collid := c.Param("collid")
		itemid := c.Param("itemid")
		if item, err := db.GetItem(collid, itemid); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		} else {
			c.JSON(http.StatusOK, item)
		}
	}
}

/**
Gets features
*/
func getFeatures(db *model.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		collectionName := c.Param("collid")

		getFeature := ogc.GetFeatureRequest{Extent: ogc.NewBbox(-180, 90, 180, -90), FeatureId: "", CollectionName: collectionName}
		//features, err := db.GetFeatures(getFeature)
		//if err != nil{
		//
		//	c.JSON(500, ogc.Exception{"500","Error fetching features"})
		//}

		fc, err := db.GetFeatures(ogc.GetFeatureRequest{CollectionName: collectionName})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"type": "FeatureCollection", "features": fc})
	}

}
