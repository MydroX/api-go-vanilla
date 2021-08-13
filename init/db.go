package init

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Initialize database connection and return pointer of this database connection
func DB() *gorm.DB {
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

	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8&parseTime=True&loc=Local"
	fmt.Println(dsn)

	var db *gorm.DB
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatal("Cannot connect to database")
	}

	return db
}
