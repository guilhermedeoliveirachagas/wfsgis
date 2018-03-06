package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/boundlessgeo/wt/handlers"
	"github.com/boundlessgeo/wt/model"
	"github.com/gin-gonic/gin"
)

type HTTPServer struct {
	server *http.Server
	router *gin.Engine
}

func NewHTTPServer() *HTTPServer {
	router := gin.Default()
	httpServer := &HTTPServer{server: &http.Server{
		Addr:    ":8080",
		Handler: router,
	}, router: router}

	r := gin.Default()
	r.Use(gin.Logger())
	// don't let errors kill the server
	r.Use(gin.Recovery())


	//handlers
	conformance := handlers.ConformanceHandler{}
	content := handlers.ContentHandler{}
	feature := handlers.FeatureHandler{}

	//Conformance endpoint
	r.GET("/api/conformance", conformance.Handle)

	//Content endpoint
	r.GET("/", content.Handle)

	//because the list of collections is dynamic we need to support random stuff here
	//any GETs that don't correspond to a feature table will 404 or throw an OGC exception
	//any POSTS should create a table
	 r.NoRoute(feature.Handle)

	return httpServer
}

//StartServer the main HTTP Server entry
func (s *HTTPServer) Start() error {
	log.Print("Starting HTTP Server")
	if err := s.server.ListenAndServe(); err != nil {
		log.Print("Error Starting Server:")
		log.Println(err)
		return err
	}
	return nil
}

//Stop for the run group
func (s *HTTPServer) Stop(err error) {
	if err != nil {
		log.Println(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	log.Print("Stopping HTTP Server")

	err = s.server.Shutdown(ctx)
	if err != nil {
		log.Println(err)
	}
}

func main() {

	db := model.NewDB("wfsthree", "wfsthree", "wfsthree", false)
	h := NewHTTPServer()
	var dbErr, httpErr error

	go func() {
		dbErr = db.Start()

		if dbErr != nil {
			log.Panic(dbErr)
		}

		httpErr = h.Start()
		if httpErr != nil {
			log.Panic(httpErr)
		}
	}()

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	running := true
	for running == true {
		select {
		case sig := <-sigchan:
			db.Stop(dbErr)
			h.Stop(httpErr)
			log.Printf("Caught signal %v\n", sig)
			running = false
		}
	}

}
