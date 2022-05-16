package model

import (
	"gorm.io/gorm"
)

type Date struct {
	gorm.Model
	AppointmentID uint   `gorm:"not null" param:"appointmentID" json:"appointmentID"`
	Date          string `gorm:"not null" json:"date"`
	// Times         Time   `gorm:"foreignKey:AppointmentID" json:"times"`
}

type PostDateResponse struct {
	DateID uint `json:"dateID"`
}
