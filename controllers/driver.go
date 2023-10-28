package controllers

import (
	"strconv"
	"taxiapp/database"
	"taxiapp/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

var validate = validator.New()

// add Driver
func AddDriver(c *gin.Context) {
	var driver models.Driver
	if err := c.Bind(&driver); err != nil {
		c.JSON(400, gin.H{
			"error": "failed to get data",
		})
		return
	}
	if err := validate.Struct(driver); err != nil {
		c.JSON(401, gin.H{
			"error": err,
		})
		return
	}

	if err := database.DB.Create(&driver).Error; err != nil {
		c.JSON(500, gin.H{
			"error": "failed to add detailes in database",
		})
		return
	}
	c.JSON(200, driver.ID)
}

// Add User Documents
func AddDocuments(c *gin.Context) {
	var documnents models.DriverDocuments
	if err := c.Bind(&documnents); err != nil {
		c.JSON(400, gin.H{
			"error": "failed to get data",
		})
		return
	}

	if err := database.DB.Create(&documnents).Error; err != nil {
		c.JSON(500, gin.H{
			"error": "failed to add detailes in database",
		})
		return
	}
	c.JSON(200, gin.H{
		"success": "successfully added driver Documents",
	})
}

// Add Vehicle Details
func AddVehicleDetails(c *gin.Context) {
	var VehicleDetails models.VehicleDetails
	if err := c.Bind(&VehicleDetails); err != nil {
		c.JSON(400, gin.H{
			"error": "failed to get data",
		})
		return
	}

	if err := database.DB.Create(&VehicleDetails).Error; err != nil {
		c.JSON(500, gin.H{
			"error": "failed to add detailes in database",
		})
		return
	}
	c.JSON(200, gin.H{
		"success": "successfully added Vehicle Details",
	})
}

// Get Driver Details
func GetDriverDetail(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("driver_id"))
	var driver models.Driver
	if err := database.DB.First(&driver, id).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "failed to find user",
		})
		return
	}
	c.JSON(200, gin.H{
		"driver": driver,
	})
}
