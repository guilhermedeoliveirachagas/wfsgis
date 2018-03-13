package model

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/boundlessgeo/wfs3/ogc"
	"github.com/lib/pq"
)

const (
	point    = iota
	mpoint   = iota
	line     = iota
	mline    = iota
	poly     = iota
	mpoly    = iota
	feat     = iota
	featcoll = iota
)

type CollectionInfoDB struct {
	geomType       int
	CollectionInfo *ogc.CollectionInfo
}

func (db *DB) AllCollectionInfos() []*CollectionInfoDB {
	//TODO impl link and extents
	qry := "SELECT table_name,description,title,crs,geom_type FROM collection_info"
	rows, err := db.db.Query(qry)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	colls := make([]*CollectionInfoDB, 0)
	for rows.Next() {
		cidb := new(CollectionInfoDB)
		ci := new(ogc.CollectionInfo)
		rows.Scan(&ci.Name, &ci.Description, &ci.Title, pq.Array(&ci.CRS), &cidb.geomType)
		cidb.CollectionInfo = ci
		colls = append(colls, cidb)
	}
	return colls
}

func (db *DB) AddCollection(coll *CollectionInfoDB) error {
	qry := "INSERT INTO collection_info (geom_type,table_name," +
		"description,title,crs) " +
		"VALUES ($1,$2,$3,$4,ARRAY[$5])"
	ci := coll.CollectionInfo
	_, err := db.db.Exec(qry, coll.geomType, ci.Name, ci.Description,
		ci.Title, strings.Join(ci.CRS, ","))

	if err != nil {
		return err
	}

	return nil
}

func (db *DB) FindCollection(collName string) *CollectionInfoDB {
	qry := "SELECT table_name,description,title,crs FROM collection_info WHERE table_name = $1"
	log.Println("Querying:" + collName)
	ci := new(ogc.CollectionInfo)
	err := db.db.QueryRow(qry, collName).Scan(&ci.Name, &ci.Description, &ci.Title, pq.Array(&ci.CRS))
	switch {
	case err == sql.ErrNoRows:
		log.Printf("No user with that ID.")
	case err != nil:
		log.Fatal(err)
	default:
		fmt.Printf("Layer is %s\n", collName)
	}
	cidb := new(CollectionInfoDB)
	cidb.CollectionInfo = ci
	return cidb
}
