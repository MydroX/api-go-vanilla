package delivery

import (
	"github.com/MydroX/api-go/internal/entities/medal"
	"github.com/gorilla/mux"
)

func MedalRoutes(r *mux.Router, h medal.Handlers) {
	r.HandleFunc("/create", h.Create).Methods("POST")
}
