package delivery

import (
	"context"
	"fmt"
	"net/http"

	"github.com/MydroX/api-go/internal/entities/medal"
	"github.com/MydroX/api-go/internal/models"
	customContext "github.com/MydroX/api-go/pkg/context"
	"github.com/MydroX/api-go/pkg/delivery"
	"github.com/gorilla/mux"
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
		return
	}
	delivery.JSONResponse(w, http.StatusCreated, "Medal created succesfully")
}

func (h *MedalHandlers) Get(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	param := mux.Vars(r)
	id := param["id"]

	medal, err := h.medalUC.Get(&ctx, id)
	if err != nil {
		if ctx.Value(customContext.HttpCode) == http.StatusNotFound {
			delivery.JSONError(w, http.StatusNotFound, fmt.Sprintf("%v", err))
			return
		}
		delivery.JSONError(w, http.StatusInternalServerError, fmt.Sprintf("Unable to get medals: %v", err))
		return
	}
	delivery.JSONResponseWithBody(w, medal, http.StatusOK, "Medal retrieved successfully")

}

func (h *MedalHandlers) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	medals, err := h.medalUC.GetAll(ctx)
	if err != nil {
		delivery.JSONError(w, http.StatusInternalServerError, fmt.Sprintf("Unable to get medals: %v", err))
		return
	}
	delivery.JSONResponseWithBody(w, medals, http.StatusOK, "Medals retrieved successfully")
}

func (h *MedalHandlers) Update(w http.ResponseWriter, r *http.Request) {

}

func (h *MedalHandlers) Delete(w http.ResponseWriter, r *http.Request) {

}
