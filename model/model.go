package model

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
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

type execSql struct {
	err error
	db  *DB
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

	if err = d.createCollectionInfoTable(); err != nil {
		log.Panic("Error creating collection info table")
	}

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

func (d *DB) createCollectionInfoTable() error {
	qry := "CREATE TABLE IF NOT EXISTS collection_info (" +
		"geom_type INTEGER," +
		"table_name TEXT," +
		"title TEXT," +
		"description TEXT," +
		"links TEXT[]," +
		"extent NUMERIC[]," +
		"crs TEXT[])"
	_, err := d.db.Exec(qry)
	if err != nil {
		return err
	}
	return nil
}
