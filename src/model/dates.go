package model

import (
	"time"

	"gorm.io/gorm"
)

type Date struct {
	gorm.Model
	AppointmentID uint      `gorm:"not null" json:"appointment_id"`
	Date          time.Time `gorm:"not null" json:"date"`
	Times         Time      `gorm:"foreignKey:AppointmentID" json:"times"`
}
