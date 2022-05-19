package model

import "gorm.io/gorm"

type URL struct {
	gorm.Model
	AppointmentID uint   `gorm:"not null" param:"appointmentID" json:"appointmentID"`
	URL           string `json:"url" validate:"required"`
}

type PostURLResponse struct {
	AppointmentURL string `json:"appointmentURL"`
}

type GetURLPayload struct {
	AppointmentID uint `param:"appointmentID"`
}
