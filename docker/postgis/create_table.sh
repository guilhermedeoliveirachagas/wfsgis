psql -U wfsthree -d wfsthree -c \
	"CREATE TABLE IF NOT EXISTS wfst_contents (
		table_name TEXT,
	 	ident TEXT,
	 	name TEXT,
	 	description TEXT,
		min_x DOUBLE PRECISION,
  		min_y DOUBLE PRECISION,
  		max_x DOUBLE PRECISION,
  		max_y DOUBLE PRECISION,
  		srs_id INTEGER)" \
	-h localhost
