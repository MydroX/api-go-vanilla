package delivery

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/MydroX/api-go/internal/entities/medal"
	"github.com/MydroX/api-go/internal/models"
	customContext "github.com/MydroX/api-go/pkg/context"
	"github.com/MydroX/api-go/pkg/delivery"
	"github.com/gorilla/mux"
)

// MedalHandlers store a medal usecase interface
type MedalHandlers struct {
	medalUC medal.UseCase
}

// NewMedalHandlers creates a new medal handlers
func NewMedalHandlers(medalUC medal.UseCase) *MedalHandlers {
	return &MedalHandlers{medalUC: medalUC}
}

// Create is a handler for creating a medal
func (h *MedalHandlers) Create(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	medal := &models.Medal{
		Name: r.FormValue("name"),
	}

	err := h.medalUC.Create(&ctx, medal)
	if err != nil {
		delivery.JSONResponse(w, http.StatusInternalServerError, fmt.Sprintf("unable to create medal: %v", err))
		return
	}
	delivery.JSONResponse(w, http.StatusCreated, "medal created succesfully")
}

// GetByID is a handler for getting a medal by id
func (h *MedalHandlers) GetByID(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	param := mux.Vars(r)
	idStr := param["id"]

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		delivery.JSONResponse(w, http.StatusInternalServerError, fmt.Sprintf("unable to get medal: %v", err))
		return
	}

	medal, err := h.medalUC.GetByID(&ctx, id)
	if err != nil {
		if ctx.Value(customContext.HTTPCode) == http.StatusNotFound {
			delivery.JSONResponse(w, http.StatusNotFound, fmt.Sprintf("%v", err))
			return
		}
		delivery.JSONResponse(w, http.StatusInternalServerError, "unable to get medals: no medal found")
		return
	}
	delivery.JSONResponseWithBody(w, medal, http.StatusOK, "medal retrieved successfully")
}

// GetAll is a handler for getting all medals
func (h *MedalHandlers) GetAll(w http.ResponseWriter, _ *http.Request) {
	ctx := context.Background()

	medals, err := h.medalUC.GetAll(&ctx)
	if err != nil {
		delivery.JSONResponse(w, http.StatusInternalServerError, fmt.Sprintf("unable to get medals: %v", err))
		return
	}
	delivery.JSONResponseWithBody(w, medals, http.StatusOK, "medals retrieved successfully")
}

// Update is a handler for updating a medal
func (h *MedalHandlers) Update(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	param := mux.Vars(r)
	idStr := param["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		delivery.JSONResponse(w, http.StatusInternalServerError, fmt.Sprintf("unable to update medal: %v", err))
		return
	}

	medal := models.Medal{
		ID:   id,
		Name: r.FormValue("name"),
	}

	err = h.medalUC.Update(&ctx, &medal)
	if err != nil {
		if ctx.Value(customContext.HTTPCode) == http.StatusNotFound {
			delivery.JSONResponse(w, http.StatusNotFound, fmt.Sprintf("unable to update medal: %v", err))
			return
		}
		delivery.JSONResponse(w, http.StatusInternalServerError, fmt.Sprintf("unable to update medal: %v", err))
		return
	}
	delivery.JSONResponse(w, http.StatusOK, "medal updated successfully")
}

// Delete is a handler for deleting a medal
func (h *MedalHandlers) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	param := mux.Vars(r)
	idStr := param["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		if ctx.Value(customContext.HTTPCode) == http.StatusNotFound {
			delivery.JSONResponse(w, http.StatusNotFound, fmt.Sprintf("unable to update medal: %v", err))
			return
		}
		delivery.JSONResponse(w, http.StatusInternalServerError, fmt.Sprintf("unable to delete medal: %v", err))
		return
	}

	err = h.medalUC.Delete(&ctx, id)
	if err != nil {
		if ctx.Value(customContext.HTTPCode) == http.StatusNotFound {
			delivery.JSONResponse(w, http.StatusNotFound, fmt.Sprintf("%v", err))
			return
		}
		delivery.JSONResponse(w, http.StatusInternalServerError, "unable to delete medals: no medal found")
		return
	}
	delivery.JSONResponse(w, http.StatusOK, "medal deleted successfully")
}
