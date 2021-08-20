package server

import (
	"net/http"

	medalDelivery "github.com/MydroX/api-go/internal/entities/medal/delivery"
	medalRepository "github.com/MydroX/api-go/internal/entities/medal/repository"
	medalUC "github.com/MydroX/api-go/internal/entities/medal/usecase"
)

func (s *Server) handlers() error {
	//Init repositories
	medalRepo := medalRepository.NewMedalRepository(s.DB)

	//Init use cases
	medalUseCase := medalUC.NewMedalUseCase(medalRepo)

	//Init handlers
	medalHandler := medalDelivery.NewMedalHandlers(medalUseCase)

	medalRouter := s.Router.PathPrefix("/medal/").Subrouter()
	medalDelivery.MedalRoutes(medalRouter, medalHandler)

	//Health check
	s.Router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"status\":\"ok\"}"))
	}).Methods("GET")

	return nil
}
