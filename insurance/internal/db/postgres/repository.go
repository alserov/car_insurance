package postgres

import (
	"context"
	"github.com/alserov/car_insurance/insurance/internal/db"
	"github.com/alserov/car_insurance/insurance/internal/service/models"
	"github.com/alserov/car_insurance/insurance/internal/utils"
	"github.com/jmoiron/sqlx"
)

func NewRepository(conn *sqlx.DB) db.Repository {
	return &postgres{conn}
}

type postgres struct {
	*sqlx.DB
}

func (p postgres) UpdateInsuranceStatus(ctx context.Context, id string, status int) error {
	query := `UPDATE insurances SET status = $1 WHERE id = $2`

	_, err := p.ExecContext(ctx, query, status, id)
	if err != nil {
		return utils.NewError(err.Error(), utils.Internal)
	}

	return nil
}

func (p postgres) CreateInsuranceData(ctx context.Context, insData models.InsuranceData) error {
	query := `INSERT INTO insurances (status, active_till, id, price) VALUES($1, $2, $3, $4)`

	_, err := p.ExecContext(ctx, query, insData.Status, insData.ActiveTill, insData.ID, insData.Price)
	if err != nil {
		return utils.NewError(err.Error(), utils.Internal)
	}

	return nil
}

func (p postgres) GetInsuranceData(ctx context.Context, ownerAddr string) (models.InsuranceData, error) {
	query := `SELECT * FROM insurances WHERE id = $1`

	row := p.QueryRowx(query, ownerAddr)

	var data models.InsuranceData
	if err := row.StructScan(&data); err != nil {
		return models.InsuranceData{}, utils.NewError(err.Error(), utils.Internal)
	}

	data.ActiveTill = data.ActiveTill.UTC()

	return data, nil
}
