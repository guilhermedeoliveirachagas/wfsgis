package handlers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/boundlessgeo/wfs3/model"
	"github.com/gin-gonic/gin"

	cors "github.com/itsjamie/gin-cors"
)

type HTTPServer struct {
	server *http.Server
	router *gin.Engine
}

func NewHTTPServer(d *model.DB) *HTTPServer {
	router := gin.Default()

	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	httpServer := &HTTPServer{server: &http.Server{
		Addr:    ":8080",
		Handler: router,
	}, router: router}

	httpServer.makeConformanceHandlers()
	httpServer.makeCollectionHandlers(d)
	httpServer.makeFeatureHandlers(d)

	return httpServer
}

//Start the main HTTP Server entry
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
