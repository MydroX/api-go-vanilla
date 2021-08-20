package mariadb

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func New() (*sql.DB, error) {
	var err error
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Cannot load .env file: %v", err)
	}

	if os.Getenv("ENV") == "dev" {
		err = godotenv.Load(".env.dev")
		if err != nil {
			log.Fatalf("Cannot load .env.dev file: %v", err)
		}
	}

	dbUser := os.Getenv("MYSQL_USER")
	dbPort := os.Getenv("MYSQL_PORT")
	dbHost := os.Getenv("MYSQL_HOST")
	dbName := os.Getenv("MYSQL_DATABASE")
	dbPassword := os.Getenv("MYSQL_PASSWORD")

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&multiStatements=true",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
	)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatalf("Cannot connect to database: %v", err)
	}

	return db, nil
}
