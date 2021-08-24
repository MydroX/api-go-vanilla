package delivery

import (
	"github.com/MydroX/api-go/internal/entities/medal"
	"github.com/gorilla/mux"
)

// MedalRoutes represents all the medal routes of the application
func MedalRoutes(r *mux.Router, h medal.Handlers) {
	r.HandleFunc("/create", h.Create).Methods("POST")
	r.HandleFunc("/get/{id:[0-9]+}", h.GetByID).Methods("GET")
	r.HandleFunc("/list", h.GetAll).Methods("GET")
	r.HandleFunc("/update/{id:[0-9]+}", h.Update).Methods("PUT")
	r.HandleFunc("/delete/{id:[0-9]+}", h.Delete).Methods("DELETE")
}
