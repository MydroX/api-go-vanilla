package repository

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/MydroX/api-go/internal/entities/medal"
	"github.com/MydroX/api-go/internal/models"
	customContext "github.com/MydroX/api-go/pkg/context"
)

type medalRepo struct {
	db *sql.DB
}

func NewMedalRepository(db *sql.DB) medal.Repository {
	return &medalRepo{
		db: db,
	}
}

func (m *medalRepo) Create(ctx *context.Context, medal *models.Medal) error {
	stmt := fmt.Sprintf(`INSERT INTO medal (name) VALUES ("%v")`, medal.Name)
	_, err := m.db.QueryContext(*ctx, stmt)

	if err != nil {
		return err
	}
	return nil
}

func (m *medalRepo) Get(ctx *context.Context, id int64) (*models.Medal, error) {
	var medal models.Medal

	err := m.db.QueryRowContext(*ctx, "SELECT * FROM medal WHERE id = ?", id).Scan(
		&medal.ID,
		&medal.Name,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			*ctx = context.WithValue(*ctx, customContext.HttpCode, http.StatusNotFound)
			return nil, err
		}
		return nil, err
	}
	return &medal, nil
}

func (m *medalRepo) GetAll(ctx *context.Context) ([]*models.Medal, error) {
	stmt := "SELECT * FROM medal"
	res, err := m.db.QueryContext(*ctx, stmt)

	var medals []*models.Medal
	for res.Next() {
		var medal models.Medal
		if err := res.Scan(&medal.ID, &medal.Name); err != nil {
			return nil, err
		}
		medals = append(medals, &medal)
	}

	if err != nil {
		return nil, err
	}
	return medals, nil
}

func (m *medalRepo) Update(ctx *context.Context, medal *models.Medal) error {
	panic("not implemented") // TODO: Implement
}

func (m *medalRepo) Delete(ctx *context.Context, id int64) error {
	panic("not implemented") // TODO: Implement
}
