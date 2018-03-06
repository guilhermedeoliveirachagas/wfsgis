package handlers

import (
	"github.com/boundlessgeo/wt/ogc"
	"github.com/boundlessgeo/wt/model"
	"encoding/json"
	"io/ioutil"
	"github.com/gin-gonic/gin"
)

type FeatureHandler struct {
	store *model.DB
}

func (fh *FeatureHandler) Handle(c *gin.Context) {

	switch c.Request.Method {
	case "GET":
		{
			collectionName := c.Request.URL.Path
			getFeature := ogc.GetFeatureRequest{Extent: ogc.NewBbox(-180, 90, 180, -90), FeatureId: "", CollectionName: collectionName}
			fh.store.GetFeatures(getFeature)

		}
	case "POST":
		{
			collectionName := c.Request.URL.Path
			fc := &ogc.FeatureCollection{}
			data, _ := ioutil.ReadAll(c.Request.Body)
			json.Unmarshal(data, fc)
			fh.store.CreateCollectionTable(collectionName, fc.Features)
		}

	default:
		{
			c.JSON(405, ogc.Exception{"405", "Method not allowed"})
		}
	}


}

