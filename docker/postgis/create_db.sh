createdb wfsthree -O wfsthree -h localhost -U postgres
psql -U postgres -d wfsthree -c "CREATE EXTENSION IF NOT EXISTS pgcrypto" -h localhost
psql -U wfsthree -d wfsthree -c "CREATE SCHEMA IF NOT EXISTS wfsthree" -h localhost
