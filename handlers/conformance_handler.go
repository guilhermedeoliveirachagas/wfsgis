package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/boundlessgeo/wt/ogc"
)

type ConformanceHandler struct {

	//stuff...

}

func(*ConformanceHandler) Handle(c *gin.Context){

	c.JSON(200,ogc.NewConformance())
}