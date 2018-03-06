package handlers

import (
	"github.com/boundlessgeo/wt/ogc"
	"github.com/gin-gonic/gin"
)

type FeatureHandler struct {
}

func (*FeatureHandler) Handle(c *gin.Context) {

	c.JSON(200, ogc.Exception{"404", "Collection doesn't exist"})
}
