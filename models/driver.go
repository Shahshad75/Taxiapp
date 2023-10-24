package models

type Input struct {
	ID          int    `json:"id" gorm:"primaryKey;unique"`
	Name        string `json:"name" gorm:"not null"`
	LastName    string `json:"last_name" gorm:"not null"`
	PhoneNumber string `json:"phone_number" gorm:"not null"`
	Email       string `json:"email" gorm:"not null;unique"`
	BirthDate   string `json:"birth_date" gorm:"not null"`
	DriverImg   string `json:"driver_img" gorm:"not null"`
	Gender      string `json:"gender" gorm:"not null"`

	CarBrand  string `json:"car_brand" gorm:"not null"`
	CarModel  string `json:"car_model" gorm:"not null"`
	CarYear   string `json:"car_year" gorm:"not null"`
	CarColor  string `json:"car_color" gorm:"not null"`
	CarSeat   string `json:"car_seat" gorm:"not null"`
	CarNumber string `json:"car_number" grom:"not null"`

	UserName string `json:"username"`
	Password string `json:"password"`

	LicenseNo    string `josn:"license_no" gorm:"not null"`
	LicenceExp   string `json:"license_exp" gorm:"not null"`
	LicenceFront string `json:"licence_ft_img" gorm:"not null"`
	LicenceBack  string `json:"licence_bk_img" gorm:"not null"`

	AdharNo      string `josn:"adhar_no" gorm:"not null"`
	AdharAddress string `json:"ahdhar_address" gorm:"not null"`
	AdharFront   string `json:"Adhar_ft_img" gorm:"not null"`
	AdharBack    string `json:"Adhar_bk_img" gorm:"not null"`

	Qualifications string `json:"qulification" gorm:"not null"`
	Expirience     string `json:"expirience" gorm:"not null"`
}

type Driver struct {
	ID             int    `json:"id" gorm:"primaryKey;unique"`
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
