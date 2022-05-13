package model

import "gorm.io/gorm"

type URL struct {
	gorm.Model
	AppointmentID uint   `gorm:"not null" json:"appointment_id"`
	URL           string `json:"url"`
}
