package model

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
	"github.com/lib/pq"

	"github.com/paulmach/orb/encoding/wkb"

	"strconv"

	"github.com/boundlessgeo/wfs3/ogc"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/encoding/wkt"
)

//creates a feature table based
func (d *DB) CreateCollectionTable(collectionName string) error {
	sql := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (_fid SERIAL PRIMARY KEY, datetime TIMESTAMP DEFAULT CURRENT_TIMESTAMP, instant TIMESTAMP, geom geometry NOT NULL, json JSONB NOT NULL)", collectionName)
	_, err := d.db.Exec(sql)
	if err != nil {
		log.Printf("Error creating table: %v", err)
		return err
	}
	return nil
}

func (d *DB) InsertFeature(collectionName string, features []*ogc.Feature) (bool, error) {
	log.Printf("INSERT FEATURE SQL: %v", features)
	insert := fmt.Sprintf("INSERT INTO %s (instant, geom, json) VALUES($1, ST_GeomFromText($2,4326), $3) RETURNING _fid as ID", collectionName)
	for _, feature := range features {
		log.Printf("FEATURE SQL: %v", feature)
		data, _ := json.Marshal(feature.Properties)
		g := wkt.MarshalString(feature.Geometry)
		var instant *time.Time
		if feature.When != nil {
			if feature.When.Type != "Instant" {
				return false, fmt.Errorf("Only 'Instant' '@type' field of 'when' is supported")
			}
			instant = feature.When.Datetime
		}
		err := d.db.QueryRow(insert, instant, g, data).Scan(&feature.ID)
		if err != nil {
			log.Printf("Error creating feature: %v", err)
			return false, err
		}
	}
	return true, nil
}

func (d *DB) GetFeatures(request ogc.GetFeatureRequest) ([]*ogc.Feature, error) {
	qry := fmt.Sprintf("SELECT _fid, instant, ST_AsBinary(geom), json FROM %s", request.CollectionName)
	rows, err := d.db.Query(qry)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	feats := make([]*ogc.Feature, 0)
	for rows.Next() {
		var id string
		sc := wkb.Scanner(nil)
		var jsonStr string
		var instant pq.NullTime
		err := rows.Scan(&id, &instant, &sc, &jsonStr)
		if err != nil {
			return nil, err
		}
		f := &ogc.Feature{ID: id}
		f.Geometry = sc.Geometry
		err = json.Unmarshal([]byte(jsonStr), &f.Properties)
		if err != nil {
			return nil, err
		}
		f.ID = id

		f.Type = sc.Geometry.GeoJSONType()

		if instant.Valid {
			f.When = &ogc.When{Type: "Instant", Datetime: &instant.Time}
		}
		feats = append(feats, f)
	}
	return feats, nil
}

/*
Delete a feature
*/
func (d *DB) DeleteItem(collectionID string, itemID string) error {

	//item id needs to be an int
	numberID, _ := strconv.Atoi(itemID)

	delete := fmt.Sprintf("DELETE from %s WHERE _fid = $1", collectionID)
	result, err := d.db.Exec(delete, numberID)
	if err != nil {
		log.Printf("Error deleting item: %v", err)
		return err
	}
	rows, err0 := result.RowsAffected()
	if err0 != nil {
		log.Printf("Error getting delete result: %v", err)
		return err
	}
	if rows == 0 {
		return fmt.Errorf("Couldn't find item %s/%s", collectionID, itemID)
	}
	log.Printf("Deleted %s/%s successfuly", collectionID, itemID)
	return nil
}

/*
Get Item by Id
*/
func (d *DB) GetItem(collectionId string, itemId string) (*ogc.FeatureCollection, error) {

	//item id needs to be an int
	numberId, _ := strconv.Atoi(itemId)

	get := fmt.Sprintf("Select _fid, ST_AsBinary(geom), json from %s WHERE _fid = $1", collectionId)

	var id int
	var g orb.Point
	var jsonStr string
	err := d.db.QueryRow(get, numberId).Scan(&id, wkb.Scanner(&g), &jsonStr)
	if err != nil {
		return nil, err
	}
	f := &ogc.Feature{ID: strconv.Itoa(id)}
	f.Geometry = g
	err = json.Unmarshal([]byte(jsonStr), &f.Properties)
	if err != nil {
		return nil, err
	}

	f.Type = "Point"

	fc := ogc.NewFeatureCollection()
	fc.Features = append(fc.Features, f)

	return fc, err

}
