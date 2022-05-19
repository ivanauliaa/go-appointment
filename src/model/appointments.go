package model

import "gorm.io/gorm"

type Appointment struct {
	gorm.Model
	UserID uint   `gorm:"not null" json:"userID"`
	Name   string `gorm:"not null" json:"name"`
	Room   string `gorm:"not null" json:"room"`
	// URL     URL    `gorm:"foreignKey:AppointmentID" json:"url"`
	// Dates   Date   `gorm:"foreignKey:AppointmentID" json:"dates"`
	// Host    User   `gorm:"foreignKey:HostID" json:"host"`
	// Guest   User   `gorm:"foreignKey:GuestID" json:"guest"`
	// Event   Event  `gorm:"foreignKey:AppointmentID"`
}

type PostAppointmentPayload struct {
	Name string `json:"name" validate:"required"`
	Room string `json:"room" validate:"required"`
}

type PostAppointmentResponse struct {
	AppointmentID uint `json:"appointmentID"`
}

type PostAppointmentConfirmPayload struct {
	AppointmentID uint `param:"appointmentID" validate:"required"`
	DateID        uint `json:"dateID" validate:"required"`
	TimeID        uint `json:"timeID" validate:"required"`
}

type AppointmentWithRelation struct {
	Appointment
	Dates []AppointmentDate `json:"dates"`
}

type GetAppointmentPayload struct {
	AppointmentID uint `param:"appointmentID"`
}
