package postgres

import (
	"github.com/alserov/car_insurance/insurance/internal/db"
	"github.com/alserov/car_insurance/insurance/internal/service/models"
)

func NewRepository() db.Repository {
	return &postgres{}
}

type postgres struct {
}

func (p postgres) UpdateInsuranceStatus() {
	//TODO implement me
	panic("implement me")
}

func (p postgres) CreateInsuranceData(insData models.InsuranceData) error {
	//TODO implement me
	panic("implement me")
}

func (p postgres) GetInsuranceData(ownerAddr string) (models.InsuranceData, error) {
	//TODO implement me
	panic("implement me")
}
