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

func (db *DB) AllCollectionInfos() []*CollectionInfoDB {
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
		colls = append(colls, c)
	}
	return colls
}

func (db *DB) AddCollection(coll *CollectionInfoDB) error {
	qry := "INSERT INTO collection_info (geom_type,table_name," +
		"description,title,extent,crs,links) " +
		"VALUES ($1,$2,$3,$4,$5,$6,$7)"

	_, err := db.db.Exec(qry, coll.geomType, coll.co.Name, coll.co.Description,
		coll.co.Title, coll.co.Extent, coll.co.CRS, coll.co.Links)

	if err != nil {
		return err
	}

	return nil
}

func (db *DB) FindCollection(collName *string) *CollectionInfoDB {
	qry := "SELECT * FROM collection_info WHERE table_name = $1"
	coll := new(CollectionInfoDB)
	db.db.QueryRow(qry).Scan(&coll)
	return coll
}
