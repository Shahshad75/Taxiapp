package routes

import (
	"taxiapp/controllers"

	"github.com/gin-gonic/gin"
)

func DriverRouter(r *gin.Engine) {
	r.POST("/adddriver", controllers.AddDriver)
	r.POST("/adddocuments", controllers.AddDocuments)
	r.POST("/addvehicledetails", controllers.AddVehicleDetails)
	r.GET("/driverdetail/:driver_id", controllers.GetDriverDetail)



}
