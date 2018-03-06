psql -U wfsthree -d wfsthree -c \
	"CREATE TABLE IF NOT EXISTS wfst_contents (
		table_name TEXT,
	 	ident TEXT,
	 	name TEXT,
	 	description TEXT,
		min_x DOUBLE,
  	min_y DOUBLE,
  	max_x DOUBLE,
  	max_y DOUBLE,
  	srs_id INTEGER)" \
	-h localhost
