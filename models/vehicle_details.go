package models

type VehicleDetails struct {
	VehicleId int    `json:"id" gorm:"primaryKey;unique"`
	DriverId  int    `json:"driverid"`
	CarBrand  string `json:"car_brand" gorm:"not null"`
	CarModel  string `json:"car_model" gorm:"not null"`
	CarYear   string `json:"car_year" gorm:"not null"`
	CarColor  string `json:"car_color" gorm:"not null"`
	CarSeat   string `json:"car_seat" gorm:"not null"`
	CarNumber string `json:"car_number" grom:"not null"`
}
