package handlers

import (
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
func updateFeature(db *model.DB) func(*gin.Context){
	return func(c *gin.Context){


	}
}
/**
Creates a feature
 */
func createFeature(db *model.DB) func(*gin.Context){
	return func(c *gin.Context){
		//collectionName := c.Param("collid")
		//fc := &ogc.FeatureCollection{}
		//data, _ := ioutil.ReadAll(c.Request.Body)
		//json.Unmarshal(data, fc)
		//db.CreateCollectionTable(collectionName, fc.Features)
	}
}
/**
Deletes a feature
 */
func deleteFeature(db *model.DB) func(*gin.Context){
	return nil
}
/**
Gets a feature by id
 */
func getFeatureById(db *model.DB) func(*gin.Context){
	return nil
}
/**
Gets features
 */
func getFeatures(db *model.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		collectionName := c.Param("collid")
		getFeature := ogc.GetFeatureRequest{Extent: ogc.NewBbox(-180, 90, 180, -90), FeatureId: "", CollectionName: collectionName}
		features, err := db.GetFeatures(getFeature)
		if err != nil{

			c.JSON(500, ogc.Exception{"500","Error fetching features"})
		}

		fc := ogc.NewFeatureCollection()
		fc.Features = features
		fc.Type = "FeatureCollection"
		c.JSON(200, fc)

	}
}



func (fh *FeatureHandler) Handle(c *gin.Context) {

	switch c.Request.Method {
	case "GET":
		{


		}
	case "POST":
		{

		}

	default:
		{
			c.JSON(405, ogc.Exception{"405", "Method not allowed"})
		}
	}

}
