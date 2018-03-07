package model

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/boundlessgeo/wt/ogc"
	"github.com/paulmach/orb/encoding/wkt"
	"strconv"
)

//creates a feature table based
func (d *DB) CreateCollectionTable(collectionName string, features []*ogc.Feature) error {

	sql := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (_fid SERIAL UNIQUE, geom geometry(Point,4326),json JSONB)", collectionName)
	_, err := d.db.Exec(sql)
	if err != nil {
		log.Printf("Error creating table: %v", err)
		return err
	}
	return nil

}

func (d *DB) InsertFeature(collectionName string, features []*ogc.Feature) (bool, error) {

	insert := fmt.Sprintf("INSERT INTO %s (geom, json) VALUES(ST_GeomFromText($1,4326), $2) RETURNING _fid as ID", collectionName)

	for _, feature := range features {

		data, _ := json.Marshal(feature)
		g := wkt.MarshalString(feature.Geometry)
		err := d.db.QueryRow(insert, g, data).Scan(&feature.ID)
		if err != nil {
			log.Printf("Error creating feature: %v", err)
			return false, err
		}
	}
	return true, nil
}

//gets features based on query
func (d *DB) GetFeatures(request ogc.GetFeatureRequest) ([]*ogc.Feature, error) {
	return nil, nil
}
/*
Delete a feature
 */
func (d *DB) DeleteItem(collectionId string, itemId string)(error){

	//item id needs to be an int
	numberId, _ := strconv.Atoi(itemId)

	delete := fmt.Sprintf("DELETE from %s WHERE _fid = $1", collectionId)
	_,err :=d.db.Exec(delete,numberId)
	if err != nil{
		log.Printf("Error deleting item: %v",err)
	}
	return err


}

/*
Get Item by Id
 */
func (d *DB) GetItem(collectionId string, itemId string)(collection *ogc.FeatureCollection,error){

	//item id needs to be an int
	numberId, _ := strconv.Atoi(itemId)

	get := fmt.Sprintf("Select fid, ST_AsWKT(geom), json from %s WHERE _fid = $1", collectionId)
	var fid int



	_,err :=d.db.QueryRow(get,numberId).Scan()
	if err != nil{
		log.Printf("Error deleting item: %v",err)
	}
	return err


}