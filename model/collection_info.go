package model

import "github.com/boundlessgeo/wt/ogc"

const (
	point = iota
	line  = iota
)

type CollectionInfoDB struct {
	geom_type int
}

func (db *DB) AllCollections() []*ogc.CollectionInfo {

	return nil
}
