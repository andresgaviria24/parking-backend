package service

import (
	"log"
	"net/http"
	"ws_parking/domain/dto"
	"ws_parking/domain/entity"
	persistence "ws_parking/infrastructure/persistence/db"
	"ws_parking/infrastructure/repository"
)

type ParkingServiceImpl struct {
	parkingRepository repository.ParkingRepository
}

func InitParkingServiceImpl() *ParkingServiceImpl {
	dbHelper, err := persistence.InitDbHelper()
	if err != nil {
		log.Fatal(err.Error())
	}
	return &ParkingServiceImpl{
		parkingRepository: dbHelper.ParkingRepository,
	}
}

func (a ParkingServiceImpl) FindParking() (dto.Response, []entity.Parking) {
	var response = dto.Response{}
	parking, err := a.parkingRepository.FindParking()
	if err != nil {
		//	response.Message = utils.Lenguage("ES", "MISSING_USER")
		//	response.Status = http.StatusNotFound
		return response, parking
	}

	response.Status = http.StatusOK
	return response, parking
}

func (a ParkingServiceImpl) CreateParking(parking entity.Parking) dto.Response {
	var response = dto.Response{}
	err := a.parkingRepository.CreateParking(parking)
	if err != nil {
		//	response.Message = utils.Lenguage("ES", "MISSING_USER")
		//	response.Status = http.StatusNotFound
		return response
	}

	response.Status = http.StatusOK
	return response
}

//func (a ParkingServiceImpl) DeleteParking()
