package persistence

import (
	"errors"
	"ws_parking/domain/entity"

	"github.com/jinzhu/gorm"
)

type ParkingRepositoryImpl struct {
	db *gorm.DB
}

func InitParkingRepositoryImpl(db *gorm.DB) *ParkingRepositoryImpl {
	return &ParkingRepositoryImpl{db}
}

func (r *ParkingRepositoryImpl) FindParking() ([]entity.Parking, error) {
	var parking []entity.Parking

	err := r.db.Find(&parking).Error

	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("Seccion no existe")
	}
	return parking, nil
}

func (r *ParkingRepositoryImpl) CreateParking(parking entity.Parking) error {
	err := r.db.Create(&parking).Error

	if err != nil {
		return err
	}
	if gorm.IsRecordNotFoundError(err) {
		return errors.New("Seccion no existe")
	}
	return nil
}
