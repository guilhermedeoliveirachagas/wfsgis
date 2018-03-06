package handlers

import (
	"github.com/boundlessgeo/wt/ogc"
	"github.com/boundlessgeo/wt/model"
	"encoding/json"
	"io/ioutil"
	"github.com/gin-gonic/gin"
	"strings"
)

type FeatureHandler struct {
	Store *model.DB
}

func (fh *FeatureHandler) Handle(c *gin.Context) {

	switch c.Request.Method {
	case "GET":
		{
			collectionName := strings.TrimLeft(c.Request.URL.Path,"/")
			getFeature := ogc.GetFeatureRequest{Extent: ogc.NewBbox(-180, 90, 180, -90), FeatureId: "", CollectionName: collectionName}
			fh.Store.GetFeatures(getFeature)

		}
	case "POST":
		{
			collectionName := strings.TrimLeft(c.Request.URL.Path,"/")
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

