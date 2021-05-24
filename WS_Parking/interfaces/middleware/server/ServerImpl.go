package server

import (
	"log"
	"os"
	"ws_parking/application"
	"ws_parking/docs"
	"ws_parking/interfaces/middleware"
	"ws_parking/utils"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files" // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger"
)

/**
Se crea un tipo de dato estandar para la API
este enruta el handler

Create standard data type for api, that on route the hadler
**/
type ServerImpl struct {
	router *gin.Engine
}

/**
Es la funcion princial de este archivo
construye las apis del servicio

Main package function, build the api service
*/
func InitServer() Server {
	prod := utils.GetBoolEnv("PROD")
	test := utils.GetBoolEnv("TEST")
	if prod == true && test == false {
		gin.SetMode(gin.ReleaseMode)
	}
	serverImpl := &ServerImpl{}
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())
	swaggerDocs()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	application.InitParkingController(router)
	serverImpl.router = router
	return serverImpl
}

func swaggerDocs() {
	docs.SwaggerInfo.Title = os.Getenv("SWAGGER_TITLE")
	docs.SwaggerInfo.Description = os.Getenv("SWAGGER_DESCRIPTION")
	docs.SwaggerInfo.Version = os.Getenv("SWAGGER_VERSION")
	docs.SwaggerInfo.Host = os.Getenv("SWAGGER_HOST")
	docs.SwaggerInfo.BasePath = os.Getenv("SWAGGER_BASEPATH")
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}

func (api *ServerImpl) RunServer() {
	appPort := os.Getenv("PORT")
	if appPort == "" {
		appPort = os.Getenv("LOCAL_PORT") //localhost
	}
	log.Fatal(api.router.Run(":" + appPort))
}
