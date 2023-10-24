package models

type Driver struct {
	DriverId       int    `json:"id" gorm:"primaryKey;unique"`
	Name           string `json:"name" gorm:"not null"`
	LastName       string `json:"last_name" gorm:"not null"`
	PhoneNumber    string `json:"phone_number" gorm:"not null"`
	Email          string `json:"email" gorm:"not null;unique"`
	BirthDate      string `json:"birth_date" gorm:"not null"`
	DriverImg      string `json:"driver_img" gorm:"not null"`
	Gender         string `json:"gender" gorm:"not null"`
	Qualifications string `json:"qulification" gorm:"not null"`
	Expirience     string `json:"expirience" gorm:"not null"`
	UserName       string `json:"username"`
	Password       string `json:"password"`

	
}
