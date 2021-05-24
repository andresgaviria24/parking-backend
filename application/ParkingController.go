package application

import (
	"net/http"
	"ws_parking/domain/entity"
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
	router.POST("/parking", parkingService.CreateParking)

}

func (a *ParkingController) GetParking(c *gin.Context) {

	response, section := a.ParkingService.FindParking()

	if response.Status != http.StatusOK {
		c.JSON(response.Status, response)
		return
	}

	c.JSON(http.StatusOK, section)
}

func (a *ParkingController) CreateParking(c *gin.Context) {

	parking := entity.Parking{}
	if err := c.ShouldBindJSON(&parking); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"invalid_json": "invalid json",
		})
		return
	}

	response := a.ParkingService.CreateParking(parking)

	if response.Status != http.StatusOK {
		c.JSON(response.Status, response)
		return
	}

	c.JSON(http.StatusOK, response)
}
