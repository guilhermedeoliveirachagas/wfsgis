package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/boundlessgeo/wt/ogc"
)

type ContentHandler struct {

	//stuff...

}

func(*ContentHandler) Handle(c *gin.Context){

	c.JSON(200,ogc.NewContent())
}