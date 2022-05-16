package model

import (
	"gorm.io/gorm"
)

type Time struct {
	gorm.Model
	DateID uint   `gorm:"not null" param:"dateID" json:"dateID"`
	Start  string `gorm:"not null" json:"start"`
	End    string `gorm:"not null" json:"end"`
}

type PostTimePayload struct {
	AppointmentID uint   `param:"appointmentID"`
	DateID        uint   `param:"dateID" json:"dateID" validate:"required"`
	Start         string `json:"start" validate:"required"`
	End           string `json:"end" validate:"required"`
}

type PostTimeResponse struct {
	TimeID uint `json:"timeID"`
}
