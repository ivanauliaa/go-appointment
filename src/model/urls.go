package model

import "gorm.io/gorm"

type URL struct {
	gorm.Model
	AppointmentID uint   `gorm:"not null" param:"appointmentID" json:"appointmentID"`
	URL           string `json:"url" validate:"required"`
}

type PostURLPayload struct {
	AppointmentID uint `param:"appointmentID"`
}

type PostURLResponse struct {
	AppointmentURL string `json:"appointmentURL"`
}
