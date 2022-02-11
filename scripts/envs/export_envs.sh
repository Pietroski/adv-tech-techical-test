#!/usr/bin/env sh

# Load variables from .env file.
# shellcheck disable=SC2046
# shellcheck disable=SC2002
export $(cat ./.env | grep -v ^# | xargs) >/dev/null

#export MIGRATIONS_PATH=internal/services/datastore/postgreSQL/${CONTEXT}/migrations/
#export PROJECT_NAME_DB_DATA_SOURCE_NAME="postgresql://node_subscriber:lnd_node_psql@localhost:5432/lnd_nodes?sslmode=disable"

# echo "MIGRATIONS_PATH=${MIGRATIONS_PATH}"
# echo "PROJECT_NAME_DB_DATA_SOURCE_NAME=${PROJECT_NAME_DB_DATA_SOURCE_NAME}"
