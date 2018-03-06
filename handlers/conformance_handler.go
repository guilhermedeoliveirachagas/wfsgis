package handlers

import (
	"net/http"

	"github.com/boundlessgeo/wt/ogc"
	"github.com/gin-gonic/gin"
)

func (h *HTTPServer) makeConformanceHandlers() {

	conformance := ogc.NewConformance()

	h.router.GET("/api/conformance", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"conformsTo": conformance})
	})
}
