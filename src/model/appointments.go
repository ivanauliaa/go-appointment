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
	Name string `json:"name"`
	Room string `json:"room"`
}

type PostAppointmentResponse struct {
	AppointmentID uint `json:"appointmentID"`
}

type PostAppointmentConfirmPayload struct {
	AppointmentID uint `param:"appointmentID"`
	DateID        uint `json:"dateID"`
	TimeID        uint `json:"timeID"`
}
