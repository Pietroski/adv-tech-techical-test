#!/usr/bin/env bash

mockgen -package mockdbtask \
  -destination internal/services/datasource/postgreSQL/task/mock/store.go \
  github.com/Pietroski/adv-tech-techical-test/internal/services/datasource/postgreSQL/task Store
