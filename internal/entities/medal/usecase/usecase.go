package usecase

import (
	"context"

	"github.com/MydroX/api-go/internal/entities/medal"
	"github.com/MydroX/api-go/internal/models"
)

type medalUC struct {
	repo medal.Repository
}

func NewMedalUseCase(repo medal.Repository) medal.UseCase {
	return &medalUC{
		repo: repo,
	}
}

func (m *medalUC) Create(ctx context.Context, medal *models.Medal) error {
	if err := m.repo.Create(ctx, medal); err != nil {
		return err
	}
	return nil
}

func (m *medalUC) Get(ctx context.Context, id int64) (*models.Medal, error) {
	panic("not implemented") // TODO: Implement
}

func (m *medalUC) Update(ctx context.Context, medal *models.Medal) error {
	panic("not implemented") // TODO: Implement
}

func (m *medalUC) Delete(ctx context.Context, id int64) error {
	panic("not implemented") // TODO: Implement
}
