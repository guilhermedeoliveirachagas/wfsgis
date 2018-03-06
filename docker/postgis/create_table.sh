psql -U wfsthree -d wfsthree -c \
	"CREATE TABLE IF NOT EXISTS collection_info (
		geom_type INTEGER
	 	name TEXT,
		title TEXT,
		description TEXT,
		links []TEXT,
		extents []NUMERIC
	 	crs TEXT,
		)" \
	-h localhost
