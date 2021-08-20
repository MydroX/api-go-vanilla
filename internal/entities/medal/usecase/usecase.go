package usecase

import (
	"context"
	"errors"
	"strconv"

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

func (m *medalUC) Get(ctx *context.Context, idStr string) (*models.Medal, error) {
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return nil, err
	}

	if id < 0 {
		return nil, errors.New("cannot get medal: wrong ID format")
	}

	medal, err := m.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return medal, nil
}

func (m *medalUC) GetAll(ctx context.Context) ([]*models.Medal, error) {
	medals, err := m.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return medals, nil
}

func (m *medalUC) Update(ctx context.Context, medal *models.Medal) error {
	panic("not implemented") // TODO: Implement
}

func (m *medalUC) Delete(ctx context.Context, id int64) error {
	panic("not implemented") // TODO: Implement
}
