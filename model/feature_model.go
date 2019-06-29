package model

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/lib/pq"

	"github.com/paulmach/orb/encoding/wkb"
	"github.com/paulmach/orb/planar"

	"strconv"

	"github.com/flaviostutz/wfsgis/ogc"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/encoding/wkt"
)

//creates a feature table based
func (d *DB) CreateCollectionTable(collectionName string) error {
	sql := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (_fid SERIAL PRIMARY KEY, datetime TIMESTAMP DEFAULT CURRENT_TIMESTAMP, instant TIMESTAMP, geom geometry NOT NULL, json JSONB NOT NULL, size INTEGER)", collectionName)
	_, err := d.db.Exec(sql)
	if err != nil {
		log.Printf("Error creating table: %v", err)
		return err
	}
	return nil
}

func (d *DB) InsertFeature(collectionName string, features []*ogc.Feature) ([]string, error) {
	insert := fmt.Sprintf("INSERT INTO %s (instant, geom, json, size) VALUES($1, ST_GeomFromText($2,4326), $3, $4) RETURNING _fid as ID", collectionName)
	var nids []string
	for _, feature := range features {
		data, _ := json.Marshal(feature.Properties)
		g := wkt.MarshalString(feature.Geometry)
		var instant *time.Time

		//get timestamp from "time" property
		ins, ok := feature.Properties["time"]
		if ok {
			ts, err := time.Parse(time.RFC3339, ins.(string))
			if err != nil {
				return []string{}, err
			}
			instant = &ts
		}

		//get timestamp from "when" attribute
		// if feature.When != nil {
		// 	if feature.When.Type != "Instant" {
		// 		return []string{}, fmt.Errorf("Only 'Instant' '@type' field of 'when' is supported")
		// 	}
		// 	instant = feature.When.Datetime
		// }

		size := float64(1)
		if feature.Geometry.Dimensions() > 0 {
			size = planar.Length(feature.Geometry)
		}

		nid := ""
		err := d.db.QueryRow(insert, instant, g, data, int(size)).Scan(&nid)
		if err != nil {
			log.Printf("Error creating feature: %v", err)
			return []string{}, err
		}
		nids = append(nids, nid)
	}
	return nids, nil
}

//GetFeatures get features
func (d *DB) GetFeatures(collectionName string, bbox *ogc.Bbox, filterAttrs map[string]string, limit int, dateStart *time.Time, dateEnd *time.Time) ([]*ogc.Feature, error) {

	w := ""
	and := ""
	if bbox != nil || len(filterAttrs) > 0 || dateStart != nil || dateEnd != nil {
		w = "WHERE "
	}
	if bbox != nil {
		w = w + fmt.Sprintf("geom && ST_MakeEnvelope(%f, %f, %f, %f, 4326)", bbox.Coords[0], bbox.Coords[1], bbox.Coords[2], bbox.Coords[3])
		and = "AND"
	}
	for k, v := range filterAttrs {
		w = w + fmt.Sprintf(" %s json->>'%s' = '%s'", and, k, v)
		and = "AND"
	}
	if dateStart != nil {
		ds := dateStart.Format(time.RFC3339)
		w = w + fmt.Sprintf(" %s instant >= '%s'", and, ds)
		and = "AND"
	}
	if dateEnd != nil {
		de := dateEnd.Format(time.RFC3339)
		w = w + fmt.Sprintf(" %s instant <= '%s'", and, de)
		and = "AND"
	}

	qry := fmt.Sprintf("SELECT _fid, instant, ST_AsBinary(geom), json FROM %s %s ORDER BY size DESC LIMIT %d;", collectionName, w, limit)
	log.Printf("GetFeatures: %s", qry)
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

		f.Type = "Feature"

		if instant.Valid {
			//add "time" property
			f.Properties["time"] = instant.Time.Format(time.RFC3339)

			//add "when" attribute
			// f.When = &ogc.When{Type: "Instant", Datetime: &instant.Time}
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

	get := fmt.Sprintf("Select _fid, instant, ST_AsBinary(geom), json from %s WHERE _fid = $1", collectionId)

	var id int
	var g orb.Point
	var jsonStr string
	sc := wkb.Scanner(&g)
	var instant pq.NullTime
	err := d.db.QueryRow(get, numberId).Scan(&id, &instant, &sc, &jsonStr)
	if err != nil {
		return nil, err
	}
	f := &ogc.Feature{ID: strconv.Itoa(id)}
	f.Geometry = g
	err = json.Unmarshal([]byte(jsonStr), &f.Properties)
	if err != nil {
		return nil, err
	}

	f.Type = "Feature"

	if instant.Valid {
		//add "time" property
		f.Properties["time"] = instant.Time.Format(time.RFC3339)

		//add "when" date info
		// f.When = &ogc.When{Type: "Instant", Datetime: &instant.Time}
	}

	fc := ogc.NewFeatureCollection()
	fc.Features = append(fc.Features, f)

	return fc, err

}
