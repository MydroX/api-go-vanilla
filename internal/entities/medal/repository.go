package medal

import (
	"context"

	"github.com/MydroX/api-go/internal/models"
)

type Repository interface {
	Create(ctx context.Context, medal *models.Medal) error
	Get(ctx context.Context, id int64) (*models.Medal, error)
	Update(ctx context.Context, medal *models.Medal) error
	Delete(ctx context.Context, id int64) error
}
