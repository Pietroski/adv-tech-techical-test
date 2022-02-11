#!/bin/sh

set -e
if [ -z "${DB_DATA_SOURCE_NAME}" ]; then
    echo "Missed PostgreSQL connection variables. Exiting ..."
    exit 1
fi
if [ -z "${MIGRATIONS_PATH}" ]; then
    echo "Missed migration path variable. Exiting ..."
    exit 1
fi

until psql "${DB_DATA_SOURCE_NAME}" -c "select 1" /dev/null 2>&1; do
    echo "Still waiting for PostgreSQL startup ..."
    sleep 1
done

echo "Starting migrations"
migrate -path "${MIGRATIONS_PATH}" -database "${DB_DATA_SOURCE_NAME}" -verbose up

echo "Starting service"
exec /build/main
