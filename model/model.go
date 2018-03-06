package model

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

//DB holds the DB connection
type DB struct {
	db      *sql.DB
	connStr string
}

//NewDB instantiates the DB struct using the injected config
func NewDB(dbname, user, pass string, sslMode bool) *DB {
	var mode string
	if sslMode {
		mode = "enable"
	} else {
		mode = "disable"
	}
	connStr := fmt.Sprintf("dbname=%s user=%s password=%s sslmode=%s",
		dbname, user, pass, mode)

	return &DB{nil, connStr}
}

//Start makes the db connection active
func (d *DB) Start() error {
	log.Println("Starting DB")
	db, err := sql.Open("postgres", d.connStr)

	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		return err
	}

	d.db = db

	log.Println("Started DB")
	return nil

}

//Stop tears down the db instance
func (d *DB) Stop(err error) {
	log.Println("Stopping DB")
	if err != nil {
		log.Println(err)
	}

	if err = d.db.Close(); err != nil {
		log.Println(err)
	}
}
