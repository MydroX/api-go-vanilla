package medal

import (
	"context"

	"github.com/MydroX/api-go/internal/models"
)

// UseCase interface for medal usecases
type UseCase interface {
	Create(ctx *context.Context, medal *models.Medal) error
	GetByID(ctx *context.Context, id int64) (*models.Medal, error)
	GetAll(ctx *context.Context) ([]*models.Medal, error)
	Update(ctx *context.Context, medal *models.Medal) error
	Delete(ctx *context.Context, id int64) error
}
