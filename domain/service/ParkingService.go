package service

import (
	"ws_parking/domain/dto"
	"ws_parking/domain/entity"
)

type ParkingService interface {
	FindParking() (dto.Response, []entity.Parking)
}
