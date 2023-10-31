package main

import (
	"fmt"
	"taxiapp/database"
	"taxiapp/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println(err)
		return
	}
	database.ConnectToDatabase()
}

func main() {
	router := gin.Default()
	routes.DriverRouter(router)
	routes.AdminRouter(router)
	router.Run(":8080")
}
