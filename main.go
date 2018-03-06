package main

import (
	"github.com/gin-gonic/gin"
	"github.com/boundlessgeo/feshack/ogc"
	"github.com/paulmach/orb"
)

func main() {
	r := gin.Default()
	r.GET("/test/wfs", func(c *gin.Context) {

		fc := ogc.NewFeatureCollection()
		p := orb.Point{0,0}
		feat := ogc.NewFeature(&p)
		feat.ID = "testid"
		feat.Properties["test"] = 1
		fc.Features = append(fc.Features, feat)

		c.JSON(200, fc)


		})

	r.Run() // listen and serve on 0.0.0.0:8080
}

