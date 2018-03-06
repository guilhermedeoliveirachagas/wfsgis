package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/boundlessgeo/wt/handlers"
	"github.com/boundlessgeo/wt/model"
	"github.com/gin-gonic/gin"
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

	//handlers
	conformance := handlers.ConformanceHandler{}
	content := handlers.ContentHandler{}
	feature := handlers.FeatureHandler{}

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	log.Printf("Starting Web Server")

	//the base endpoint should provide a list of all the supported collections
	// aka tables
	r.GET("/collection/:collectionId/", feature.Handle)

	//Conformance endpoint
	r.GET("/api/conformance", conformance.Handle)

	//Content endpoint
	r.GET("/", content.Handle)

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
