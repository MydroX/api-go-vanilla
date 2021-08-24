package server

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

// Server stores the router and the database connection of the application
type Server struct {
	Router *mux.Router
	DB     *sql.DB
}

// NewServer create a new server
func NewServer(db *sql.DB) *Server {
	return &Server{
		DB: db,
	}
}

// Run the server of the application
func (s *Server) Run() error {
	s.Router = mux.NewRouter()

	if err := s.handlers(); err != nil {
		return err
	}

	err := http.ListenAndServe(":8080", s.Router)
	if err != nil {
		return err
	}

	return nil
}
