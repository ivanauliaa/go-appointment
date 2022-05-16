package model

import (
	"gorm.io/gorm"
)

type Date struct {
	gorm.Model
	AppointmentID uint   `gorm:"not null" param:"appointmentID" json:"appointmentID" validate:"required"`
	Date          string `gorm:"not null" json:"date" validate:"required"`
	// Times         Time   `gorm:"foreignKey:AppointmentID" json:"times"`
}

type PostDateResponse struct {
	DateID uint `json:"dateID"`
}
