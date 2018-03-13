package http

import (
	"context"
	"log"
	"net/http"
	"time"

	_ "github.com/boundlessgeo/wt/handlers"
	"github.com/boundlessgeo/wt/model"
	"github.com/gin-gonic/gin"
)

//HTTPServer holds references to the Server and the router.
//This can be used to provide a lifecycle for the web server
type HTTPServer struct {
	server *http.Server
	router *gin.Engine
}

func NewHTTPServer(d *model.DB) *HTTPServer {
	router := gin.Default()
	httpServer := &HTTPServer{server: &http.Server{
		Addr:    ":8080",
		Handler: router,
	}, router: router}

	httpServer.MakeConformanceHandlers()
	httpServer.MakeCollectionHandlers(d)
	httpServer.MakeFeatureHandlers(d)

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
