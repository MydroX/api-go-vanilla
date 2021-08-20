package server

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	Router *mux.Router
	DB     *sql.DB
}

func NewServer(db *sql.DB) *Server {
	return &Server{
		DB: db,
	}
}

func (s *Server) Run() error {
	s.Router = mux.NewRouter()

	if err := s.handlers(); err != nil {
		return err
	}

	if err := http.ListenAndServe(":8080", s.Router); err != nil {
		return err
	}

	return nil
}
