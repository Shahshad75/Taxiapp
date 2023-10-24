package routes

import (
	"taxiapp/controllers"

	"github.com/gin-gonic/gin"
)

func DriverRouter(router *gin.Engine) {
	router.POST("/createdriver",controllers.CreateDriver)
	router.GET("/driverdetail/:driver_id",controllers.GetDriverDetail)
}
