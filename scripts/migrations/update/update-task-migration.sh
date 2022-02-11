#!/usr/bin/env bash

migration_name=${1?"migration name required"}

migrate create -ext sql -dir internal/services/datasource/postgreSQL/task/migrations -seq $migration_name
