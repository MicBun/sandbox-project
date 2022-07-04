package core

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string
}

type Order struct {
	gorm.Model
	TrackingNumber      string `gorm:"unique"`
	ConsigneeAddress    string
	ConsigneeCity       string
	ConsigneeProvince   string
	ConsigneePostalCode string
	ConsigneeCountry    string
	Weight              float32
	Height              float32
	Width               float32
	Length              float32
	UserID              uint `gorm:"not_null"`
	User                User
}

type ServiceProvider struct {
	gorm.Model
	Name string
}
