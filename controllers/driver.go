package controllers

import (
	"strconv"
	"taxiapp/database"
	"taxiapp/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validat = validator.New()

func CreateDriver(c *gin.Context) {
	var driver models.Input
	if err := c.Bind(&driver); err != nil {
		c.JSON(400, gin.H{
			"error": "failed to get data",
		})
		return
	}
	if err := validat.Struct(driver); err != nil {
		c.JSON(401, gin.H{
			"error": err,
		})
		return
	}

	if err := database.DB.Create(&models.Driver{
		Name:           driver.Name,
		LastName:       driver.LastName,
		PhoneNumber:    driver.PhoneNumber,
		Email:          driver.Email,
		BirthDate:      driver.BirthDate,
		DriverImg:      driver.DriverImg,
		Gender:         driver.Gender,
		Qualifications: driver.Qualifications,
		Expirience:     driver.Expirience,
	}).Error; err != nil {
		c.JSON(500, gin.H{
			"error": "failed to add DriverDetails in database",
		})
		return
	}

	if err := database.DB.Create(&models.VehicleDetails{
		DriverId:  driver.ID,
		CarBrand:  driver.CarBrand,
		CarModel:  driver.CarModel,
		CarYear:   driver.CarYear,
		CarColor:  driver.CarColor,
		CarSeat:   driver.CarSeat,
		CarNumber: driver.CarNumber,
	}).Error; err != nil {
		c.JSON(500, gin.H{
			"error": "failed to add VehicleDetails in database",
		})
		return
	}

	if err := database.DB.Create(&models.DriverDocument{
		DriverId:     driver.ID,
		LicenseNo:    driver.LicenseNo,
		LicenceExp:   driver.LicenceExp,
		LicenceFront: driver.LicenceFront,
		LicenceBack:  driver.LicenceBack,
		AdharNo:      driver.AdharNo,
		AdharAddress: driver.AdharAddress,
		AdharFront:   driver.AdharFront,
		AdharBack:    driver.AdharBack,
	}).Error; err != nil {
		c.JSON(500, gin.H{
			"error": "failed to add DriverDocumets in database",
		})
		return
	}

	c.JSON(200, gin.H{
		"success": "successfully created driver" + driver.Name,
	})
}

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
