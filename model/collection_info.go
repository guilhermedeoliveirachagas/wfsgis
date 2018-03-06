package model

const (
	point = iota
	line  = iota
)

type CollectionInfoDB struct {
	geom_type int
}

func (db *DB) AllCollections() []*CollectionInfo {
	qry := "SELECT"
	rows, err := db.db.Query(qry)
	if err != nil {

	}
	return nil
}
