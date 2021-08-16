package server

import (
	"database/sql"
)

type Server struct {
	DB *sql.DB
}

func NewServer(db *sql.DB) *Server {
	return &Server{
		DB: db,
	}
}

func (s *Server) Run() error {
	if err := handlers(); err != nil {
		return err
	}

	return nil
}
