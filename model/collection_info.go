package model

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/flaviostutz/wfsgis/ogc"
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

func (db *DB) AddCollection(coll *CollectionInfoDB) (bool, error) {
	tx, err := db.db.Begin()
	if err != nil {
		return false, fmt.Errorf("Error on starting add collection transaction: %s", err)
	}

	dmlQuery := "INSERT INTO collection_info (geom_type,table_name," +
		"description,title,crs) " +
		"VALUES ($1,$2,$3,$4,ARRAY[$5])"
	ci := coll.CollectionInfo

	if _, insErr := tx.Exec(dmlQuery, coll.geomType, ci.Name, ci.Description,
		ci.Title, strings.Join(ci.CRS, ",")); insErr != nil {
		tx.Rollback()
		pqErr := insErr.(*pq.Error)
		if pqErr.Code == "23505" {
			return true, fmt.Errorf("Collection named {%s} already created", coll.CollectionInfo.Name)
		}
		return false, fmt.Errorf("Error %s while inserting metadata for new collection named {%s}", insErr, coll.CollectionInfo.Name)
	}

	ddlQuery := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (_fid SERIAL PRIMARY KEY, datetime TIMESTAMP DEFAULT CURRENT_TIMESTAMP, instant TIMESTAMP, geom geometry NOT NULL, json JSONB NOT NULL, size INTEGER)", coll.CollectionInfo.Name)
	if _, creErr := tx.Exec(ddlQuery); creErr != nil {
		tx.Rollback()
		return false, fmt.Errorf("Error %s while creating table for collection named {%s}", creErr, coll.CollectionInfo.Name)
	}

	tx.Commit()
	return false, nil
}

func (db *DB) RemoveCollection(collName string) (bool, error) {

	tx, err := db.db.Begin()
	if err != nil {
		return false, fmt.Errorf("Error on starting remove collection transaction: %s", err)
	}

	dmlQuery := "DELETE from collection_info WHERE table_name = $1"

	if res, delErr := tx.Exec(dmlQuery, collName); delErr != nil {
		tx.Rollback()
		return false, fmt.Errorf("Error %s while deleting table %s", delErr, collName)
	} else if rowsAffected, cntErr := res.RowsAffected(); cntErr != nil {
		tx.Rollback()
		return false, fmt.Errorf("Error %s reading delete result", delErr, collName)
	} else if rowsAffected != 1 {
		tx.Rollback()
		return true, nil
	}

	ddlQuery := fmt.Sprintf("DROP TABLE %s", collName)
	_, delTblErr := tx.Exec(ddlQuery)
	if delTblErr != nil {
		tx.Rollback()
		pqErr := delTblErr.(*pq.Error)
		if pqErr.Code == "42P01" {
			return true, fmt.Errorf("Collection %s is on inconsistent state", collName)
		}
		return false, fmt.Errorf("Error dropping table %s, reason %s", collName, delTblErr)
	}

	tx.Commit()
	return false, nil
}

func (db *DB) FindCollection(collName string) (*CollectionInfoDB, error) {
	qry := "SELECT table_name,description,title,crs FROM collection_info WHERE table_name = $1"
	log.Println("Querying:" + collName)
	ci := new(ogc.CollectionInfo)
	err := db.db.QueryRow(qry, collName).Scan(&ci.Name, &ci.Description, &ci.Title, pq.Array(&ci.CRS))
	if err != nil {
		if err == sql.ErrNoRows {
			return &CollectionInfoDB{}, err
		}
		return nil, err
	}
	cidb := new(CollectionInfoDB)
	cidb.CollectionInfo = ci
	return cidb, nil
}

func (d *DB) createCollectionInfoTable() error {
	qry := "CREATE TABLE IF NOT EXISTS collection_info (" +
		"table_name TEXT PRIMARY KEY," +
		"geom_type INTEGER NOT NULL," +
		"title TEXT NOT NULL," +
		"description TEXT NOT NULL," +
		"links TEXT[]," +
		"extent NUMERIC[]," +
		"crs TEXT[] NOT NULL)"
	_, err := d.db.Exec(qry)
	if err != nil {
		return err
	}
	return nil
}
