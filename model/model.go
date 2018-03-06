package model

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq"
	"github.com/boundlessgeo/wt/ogc"
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


//creates a feature table based
func(d *DB) CreateCollectionTable(collectionName string, features []*ogc.Feature) error{

	 sql := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (_fid SERIAL UNIQUE, geom geometry(Point,4326),json JSONB)",collectionName)
	_, err := d.db.Exec(sql)
	if err != nil {
		log.Printf("Error creating table: %v",err)
		return err
	}
	//makeGeom := "SELECT ST_AddGeometryColumn('public',$1,'geom',4326,'POINT',2)"
	//_, err = d.db.Exec(makeGeom,collectionName)
	//if err != nil {
	//	log.Printf("Error adding geometry column: %v",err)
	//	return err
	//}
	return nil

}
//gets features based on query
func(d *DB) GetFeatures(request ogc.GetFeatureRequest) ([]*ogc.Feature, error) {
	return nil, nil
}
