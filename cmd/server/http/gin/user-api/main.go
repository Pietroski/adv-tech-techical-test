package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	userFactory "github.com/Pietroski/adv-tech-techical-test/internal/factories/user"
	"github.com/Pietroski/adv-tech-techical-test/internal/services/datasource/postgreSQL/user"
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
	userStore := user.NewStore(dbConn)

	userServer := userFactory.NewUserFactory(userStore)
	if err = userServer.Start("server-address"); err != nil {
		log.Fatalf("cannot start server -> %s", err.Error())
	}
}
