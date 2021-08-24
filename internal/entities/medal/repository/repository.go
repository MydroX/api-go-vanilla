package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/MydroX/api-go/internal/entities/medal"
	"github.com/MydroX/api-go/internal/models"
	mydroxContext "github.com/MydroX/api-go/pkg/context"
	mydroxTime "github.com/MydroX/api-go/pkg/time"
)

type medalRepo struct {
	db *sql.DB
}

// NewMedalRepository returns a new instance of the medal repository
func NewMedalRepository(db *sql.DB) medal.Repository {
	return &medalRepo{
		db: db,
	}
}

func (m *medalRepo) Create(ctx *context.Context, medal *models.Medal) error {
	stmt := fmt.Sprintf(`INSERT INTO medal (name, created_at, updated_at) VALUES ("%v", "%v", "%v")`, medal.Name, mydroxTime.TimeToMySQLTime(medal.CreatedAt), mydroxTime.TimeToMySQLTime(medal.UpdatedAt))
	_, err := m.db.QueryContext(*ctx, stmt)

	if err != nil {
		return err
	}
	return nil
}

func (m *medalRepo) GetByID(ctx *context.Context, id int64) (*models.Medal, error) {
	var medal models.Medal

	err := m.db.QueryRowContext(*ctx, "SELECT * FROM medal WHERE id = ?", id).Scan(
		&medal.ID,
		&medal.Name,
		&medal.CreatedAt,
		&medal.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			*ctx = context.WithValue(*ctx, mydroxContext.HTTPCode, http.StatusNotFound)
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

		err := res.Scan(
			&medal.ID,
			&medal.Name,
			&medal.CreatedAt,
			&medal.UpdatedAt,
		)

		if err != nil {
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
	res, err := m.db.ExecContext(*ctx, "UPDATE medal SET name = ?, updated_at = ? WHERE id = ?", medal.Name, mydroxTime.GetTimeNowString(), medal.ID)

	if v, _ := res.RowsAffected(); v == 0 {
		*ctx = context.WithValue(*ctx, mydroxContext.HTTPCode, http.StatusNotFound)
		return errors.New("id not found")
	}

	if err != nil {
		return err
	}

	return nil
}

func (m *medalRepo) Delete(ctx *context.Context, id int64) error {
	res, err := m.db.ExecContext(*ctx, "DELETE FROM medal WHERE id = ?", id)

	if v, _ := res.RowsAffected(); v == 0 {
		*ctx = context.WithValue(*ctx, mydroxContext.HTTPCode, http.StatusNotFound)
		return errors.New("id not found")
	}

	if err != nil {
		return err
	}
	return nil
}
