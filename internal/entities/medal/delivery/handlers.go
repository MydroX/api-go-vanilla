package delivery

import (
	"context"
	"fmt"
	"net/http"

	"github.com/MydroX/api-go/internal/entities/medal"
	"github.com/MydroX/api-go/internal/models"
	"github.com/MydroX/api-go/pkg/delivery"
)

type MedalHandlers struct {
	medalUC medal.UseCase
}

func NewMedalHandlers(medalUC medal.UseCase) *MedalHandlers {
	return &MedalHandlers{medalUC: medalUC}
}

func (h *MedalHandlers) Create(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	medal := &models.Medal{
		Name: r.FormValue("name"),
	}

	err := h.medalUC.Create(ctx, medal)
	if err != nil {
		delivery.JSONError(w, http.StatusInternalServerError, fmt.Sprintf("Unable to create medal: %v", err))
	}
	delivery.JSONResponse(w, nil, http.StatusCreated, "Medal created succesfully")
}

func (h *MedalHandlers) Get(w http.ResponseWriter, r *http.Request) {

}

func (h *MedalHandlers) Update(w http.ResponseWriter, r *http.Request) {

}

func (h *MedalHandlers) Delete(w http.ResponseWriter, r *http.Request) {

}
