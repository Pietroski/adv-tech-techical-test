#!/usr/bin/env bash

migrate create -ext sql -dir internal/services/datasource/postgreSQL/task/migrations -seq create_tasks_table
