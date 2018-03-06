package main

import (
	"github.com/gin-gonic/gin"
	"github.com/boundlessgeo/feshack/ogc"
	"github.com/paulmach/orb"
	"github.com/boundlessgeo/feshack/handlers"
)

func main() {
	r := gin.Default()


	//handlers
	conformance := handlers.ConformanceHandler{}
	content := handlers.ContentHandler{}
	feature := handlers.FeatureHandler{}

	//hacked PoC feature encoding endpoint
	r.GET("/test/wfs", func(c *gin.Context) {

		fc := ogc.NewFeatureCollection()
		p := orb.Point{0,0}
		feat := ogc.NewFeature(&p)
		feat.ID = "testid"
		feat.Properties["test"] = 1
		fc.Features = append(fc.Features, feat)

		c.JSON(200, fc)

		})

	//Conformance endpoint
	r.GET("/api/conformance",conformance.Handle)

	//Content endpoint
	r.GET("/",content.Handle)

	//the base endpoint should provide a list of all the supported collections
	// aka tables
	r.GET("/:collectionId", feature.Handle);

	r.Run() // listen and serve on 0.0.0.0:8080
}
