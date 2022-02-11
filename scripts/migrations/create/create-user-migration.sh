#!/usr/bin/env bash

migrate create -ext sql -dir internal/services/datasource/postgreSQL/user/migrations -seq create_users_table
