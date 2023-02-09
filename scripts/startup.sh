#!/bin/sh
set -e

./migrate -verbose -source file://db/migration -database "mysql://${DB_USERNAME}:${DB_PASSWORD}@tcp(${DB_HOST})/${DB_NAME}" up
exec $@
