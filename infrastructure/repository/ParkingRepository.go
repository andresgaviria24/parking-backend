package repository

import "ws_parking/domain/entity"

type ParkingRepository interface {
	FindParking() ([]entity.Parking, error)
	CreateParking(entity.Parking) error
	//DeleteParking(id int) error
	//ModifiedParking(entity.Parking) error
}
