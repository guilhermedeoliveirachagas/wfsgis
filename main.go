package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/boundlessgeo/feshack/ogc"
	"github.com/boundlessgeo/wt/model"
	"github.com/gin-gonic/gin"
	"github.com/paulmach/orb"
)

func main() {

	db := model.NewDB("wfsthree", "wfsthree", "wfsthree", false)
	var dbErr error

	go func() {
		dbErr = db.Start()

		if dbErr != nil {
			log.Panic(dbErr)
		}
	}()

	r := gin.Default()
	r.GET("/test/wfs", func(c *gin.Context) {

		fc := ogc.NewFeatureCollection()
		p := orb.Point{0, 0}
		feat := ogc.NewFeature(&p)
		feat.ID = "testid"
		feat.Properties["test"] = 1
		fc.Features = append(fc.Features, feat)

		c.JSON(200, fc)

	})

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	log.Printf("Starting Web Server")
	r.Run() // listen and serve on 0.0.0.0:8080

	running := true
	for running == true {
		select {
		case sig := <-sigchan:
			db.Stop(dbErr)
			log.Printf("Caught signal %v\n", sig)
			running = false
		}
	}

}
