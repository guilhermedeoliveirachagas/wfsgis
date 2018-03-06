createdb wfsthree -O wfsthree -h localhost -U postgres
psql -U postgres -d wfsthree -c "CREATE EXTENSION IF NOT EXISTS pgcrypto; CREATE EXTENSION IF NOt EXISTS postgis" -h localhost
