package application

import (
	"net/http"
	"ws_parking/domain/service"

	"github.com/gin-gonic/gin"
)

type ParkingController struct {
	ParkingService service.ParkingService
}

func InitParkingController(router *gin.Engine) {
	parkingService := ParkingController{
		ParkingService: service.InitParkingServiceImpl(),
	}
	router.GET("/parking", parkingService.GetParking)

}

func (a *ParkingController) GetParking(c *gin.Context) {

	response, section := a.ParkingService.FindParking()

	if response.Status != http.StatusOK {
		c.JSON(response.Status, response)
		return
	}

	c.JSON(http.StatusOK, section)
}
