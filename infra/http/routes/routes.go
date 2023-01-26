package routes

import (
	"csv-wrapper/application/controllers/echo"
	"github.com/gin-gonic/gin"
)

func AddRoutes() *gin.Engine {
	router := gin.Default()
	router.POST("/echo", controller_echo.EchoCsv)
	return router
}
