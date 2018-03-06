package model

import (
	"github.com/boundlessgeo/wt/ogc"
	"fmt"
	"log"
)

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

func(d *DB) InsertFeature(collectionName string, features []*ogc.Feature){

	//insert := "INSERT INTO $1 (geom, json) VALUES($2, $3)"
	//
	//for _,feature := range features{
	//
	//	json, _ := json2.Marshal(feature)
	//	orb.AllGeometries.


	//	}


}

//gets features based on query
func(d *DB) GetFeatures(request ogc.GetFeatureRequest) ([]*ogc.Feature, error) {
	return nil, nil
}