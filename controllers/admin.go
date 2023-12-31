package controllers

import (
	"net/http"
	"os"
	"strconv"
	"taxiapp/database"
	"taxiapp/helper"
	"taxiapp/models"

	"github.com/gin-gonic/gin"
)

// Login
func AdminLogin(c *gin.Context) {
	var credentials struct {
		Username string `json:"username"`
		Password string ` json:"password"`
	}

	if err := c.Bind(&credentials); err != nil {
		c.JSON(400, gin.H{
			"error": "failed to bind",
		})
		return
	}
	username := os.Getenv("USER_NAME")
	password := os.Getenv("PASSWORD")
	if username != credentials.Username && password != credentials.Password {
		c.JSON(400, gin.H{
			"error": "incorrect username or password",
		})
		return
	}
	token, err := helper.GenerateJWTToken(password)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to generate token",
		})
		return
	}
	//set token into browser
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("jwt_admin", token, 3600*24*30, "", "", true, true)
	//success message
	c.JSON(200, gin.H{
		"message": "Successfully logged admin",
	})

}

// get pending drivers
func GetPendingDrivers(c *gin.Context) {
	var details []models.Detail
	if err := database.DB.Table("drivers").
		Select("driver_documents.license_no,driver_documents.licence_exp,driver_documents.licence_front,driver_documents.licence_back,driver_documents.adhar_no,driver_documents.adhar_address,driver_documents.adhar_front,adhar_back,drivers.id,drivers.name,drivers.last_name,drivers.phone_number,drivers.email,drivers.birth_date,drivers.driver_img,drivers.gender,drivers.qualification,drivers.experience,drivers.status,vehicle_details.vehicle_brand,vehicle_details.vehicle_model,vehicle_details.vehicle_year,vehicle_details.vehicle_color,vehicle_details.vehicle_seat,vehicle_details.vehicle_number").
		Joins("INNER JOIN driver_documents ON driver_documents.user_id=drivers.id").
		Joins("INNER JOIN vehicle_details ON vehicle_details.user_id=drivers.id").
		Where("status=?", "Pending").Scan(&details).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "error fetching data",
		})
		return
	}
	c.JSON(200, gin.H{
		"datas": details,
	})

}

// Accept Driver
func AcceptDrivers(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var driver models.Driver
	if err := database.DB.First(&driver, id).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "failed to find user",
		})
		return
	}
	var body struct {
		Status   string `json:"status"`
		UserName string `json:"user_name"`
		Password string `json:"password"`
	}
	if err := c.Bind(&body); err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to Bind Data",
		})
	}
	// Set the email as the username and the password as the phone number
	data := map[string]interface{}{
		"status":    "Accepted",
		"user_name": driver.Email,
		"password":  driver.PhoneNumber + "@1234",
	}

	// Update the record with the new credentials
	if err := database.DB.Model(&models.Driver{}).Where("id = ?", id).Updates(data).Error; err != nil {
		c.JSON(500, gin.H{
			"error": "failed to update details in database",
		})
		return
	}
	c.JSON(200, gin.H{
		"success": "Accepted",
	})

}

//Reject Driver

func RejectDrivers(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := database.DB.Model(&models.Driver{}).Where("id=?", id).Update("status", "Rejected").Error; err != nil {
		c.JSON(400, gin.H{
			"error": "failed to find user",
		})
		return
	}
	c.JSON(200, gin.H{
		"success": "Rejected",
	})

}

// Get Accepted Drivers
func GetAcceptedDrivers(c *gin.Context) {
	var details []models.Detail
	if err := database.DB.Table("drivers").
		Select("driver_documents.license_no,driver_documents.licence_exp,driver_documents.licence_front,driver_documents.licence_back,driver_documents.adhar_no,driver_documents.adhar_address,driver_documents.adhar_front,adhar_back,drivers.id,drivers.name,drivers.last_name,drivers.phone_number,drivers.email,drivers.birth_date,drivers.driver_img,drivers.gender,drivers.qualification,drivers.experience,drivers.status,vehicle_details.vehicle_brand,vehicle_details.vehicle_model,vehicle_details.vehicle_year,vehicle_details.vehicle_color,vehicle_details.vehicle_seat,vehicle_details.vehicle_number").
		Joins("INNER JOIN driver_documents ON driver_documents.user_id=drivers.id").
		Joins("INNER JOIN vehicle_details ON vehicle_details.user_id=drivers.id").
		Where("status=?", "Accepted").Scan(&details).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "error fetching data",
		})
		return
	}
	c.JSON(200, gin.H{
		"datas": details,
	})

}

// Get Rejected Drivers
func GetRejectedDrivers(c *gin.Context) {
	var details []models.Detail
	if err := database.DB.Table("drivers").
		Select("driver_documents.license_no,driver_documents.licence_exp,driver_documents.licence_front,driver_documents.licence_back,driver_documents.adhar_no,driver_documents.adhar_address,driver_documents.adhar_front,adhar_back,drivers.id,drivers.name,drivers.last_name,drivers.phone_number,drivers.email,drivers.birth_date,drivers.driver_img,drivers.gender,drivers.qualification,drivers.experience,drivers.status,vehicle_details.vehicle_brand,vehicle_details.vehicle_model,vehicle_details.vehicle_year,vehicle_details.vehicle_color,vehicle_details.vehicle_seat,vehicle_details.vehicle_number").
		Joins("INNER JOIN driver_documents ON driver_documents.user_id=drivers.id").
		Joins("INNER JOIN vehicle_details ON vehicle_details.user_id=drivers.id").
		Where("status=?", "Rejected").Scan(&details).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "error fetching data",
		})
		return
	}
	c.JSON(200, gin.H{
		"datas": details,
	})

}
