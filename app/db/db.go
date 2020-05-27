package db

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

var (
	dbName     string
	dbUserName string
	dbPassword string
	dbHost     string
)

func InitDbConnection() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)

		switch pair[0] {
		case "PG_PASSWORD":
			dbPassword = pair[1]
		case "PG_USERNAME":
			dbUserName = pair[1]
		case "PG_DBNAME":
			dbName = pair[1]
		case "DB_HOST":
			dbHost = pair[1]
		}
	}

	connStr := fmt.Sprintf("postgres://%v:%v@%v/%v?sslmode=disable", dbUserName, dbPassword, dbHost, dbName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("error connecting to db ", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("db ping failed ", err)
	}

	return db
}
