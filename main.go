package main

import (
	"log"

	"github.com/MydroX/api-go/internal/server"
	"github.com/MydroX/api-go/pkg/db/mariadb"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	log.Println("Starting API Server...")
	db, err := mariadb.New()
	if err != nil {
		log.Fatalf("DB Error : %v", err)
	}

	s := server.NewServer(db)
	if err := s.Run(); err != nil {
		log.Fatalf("Server error : %v", err)
	}
}
