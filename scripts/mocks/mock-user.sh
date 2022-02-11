#!/usr/bin/env bash

mockgen -package mockdbuser \
  -destination internal/services/datasource/postgreSQL/user/mock/store.go \
  github.com/Pietroski/adv-tech-techical-test/internal/services/datasource/postgreSQL/user Store
