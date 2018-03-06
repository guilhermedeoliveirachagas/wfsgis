package model

import (
	"log"

	"github.com/boundlessgeo/wt/ogc"
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
	geomType int
	co       ogc.CollectionInfo
}

func (db *DB) AllCollections() []*CollectionInfoDB {
	qry := "SELECT * FROM collection_info"
	rows, err := db.db.Query(qry)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	colls := make([]*CollectionInfoDB, 0)
	for rows.Next() {
		c := new(CollectionInfoDB)
		rows.Scan(&c.co)
		rows.Scan(&c.geomType)
	}
	return colls
}
