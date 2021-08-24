package usecase

import (
	"context"
	"errors"

	"github.com/MydroX/api-go/internal/entities/medal"
	"github.com/MydroX/api-go/internal/models"
)

type medalUC struct {
	repo medal.Repository
}

// NewMedalUseCase creates a new usecase for the medal entity
func NewMedalUseCase(repo medal.Repository) medal.UseCase {
	return &medalUC{
		repo: repo,
	}
}

func (m *medalUC) Create(ctx *context.Context, medal *models.Medal) error {
	if medal.Name == "" {
		return errors.New("invalid name")
	}

	err := m.repo.Create(ctx, medal)
	if err != nil {
		return err
	}
	return nil
}

func (m *medalUC) GetByID(ctx *context.Context, id int64) (*models.Medal, error) {
	if id < 0 {
		return nil, errors.New("cannot get medal: invalid ID")
	}

	medal, err := m.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return medal, nil
}

func (m *medalUC) GetAll(ctx *context.Context) ([]*models.Medal, error) {
	medals, err := m.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return medals, nil
}

func (m *medalUC) Update(ctx *context.Context, medal *models.Medal) error {
	if medal.Name == "" {
		return errors.New("invalid name")
	}

	if medal.ID <= 0 {
		return errors.New("invalid ID")
	}

	err := m.repo.Update(ctx, medal)

	if err != nil {
		return err
	}
	return nil
}

func (m *medalUC) Delete(ctx *context.Context, id int64) error {
	if id <= 0 {
		return errors.New("invalid ID")
	}

	err := m.repo.Delete(ctx, id)

	if err != nil {
		return err
	}
	return nil
}
