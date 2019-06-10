package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/boundlessgeo/wfs3/handlers"
	"github.com/boundlessgeo/wfs3/model"
)

func main() {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	dbname := os.Getenv("POSTGRES_DBNAME")
	username := os.Getenv("POSTGRES_USERNAME")
	password := os.Getenv("POSTGRES_PASSWORD")

	db := model.NewDB(dbname, username, password, false, host, port)
	h := handlers.NewHTTPServer(db)
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
