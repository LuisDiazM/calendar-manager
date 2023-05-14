#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
	CREATE USER MsManager;
	CREATE DATABASE CalendarManager;
	GRANT ALL PRIVILEGES ON DATABASE CalendarManager TO MsManager;
EOSQL