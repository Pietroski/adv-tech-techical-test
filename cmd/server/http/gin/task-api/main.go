package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	taskFactory "github.com/Pietroski/adv-tech-techical-test/internal/factories/task"
	"github.com/Pietroski/adv-tech-techical-test/internal/services/datasource/postgreSQL/task"
)

var (
	// apiConfig models.APIStruct
	dbConn *sql.DB
	err    error
)

func init() {
	dbConn, err = sql.Open("apiConfig.DBDriver", "apiConfig.DBDataSourceName")
	if err != nil {
		log.Fatalf("Error connecting to database -> %s", err.Error())
	}
}

func main() {
	taskStore := task.NewStore(dbConn)

	userServer := taskFactory.NewTaskFactory(taskStore)
	if err = userServer.Start("server-address"); err != nil {
		log.Fatalf("cannot start server -> %s", err.Error())
	}
}
