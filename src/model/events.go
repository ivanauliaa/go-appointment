package model

import "gorm.io/gorm"

type Event struct {
	gorm.Model
	AppointmentID uint `gorm:"not null" json:"appointmentID" validate:"required"`
	HostID        uint `gorm:"not null" json:"hostID" validate:"required"`
	GuestID       uint `gorm:"not null" json:"guestID" validate:"required"`
	// Name          string `gorm:"not null" json:"name"`
	DateID uint   `json:"dateID" validate:"required"`
	TimeID uint   `json:"timeID" validate:"required"`
	Status string `gorm:"not null" json:"status" validate:"required"`
	// Appointment   Appointment `gorm:"foreignKey: AppointmentID" json:"appointment"`
	// Date Date `gorm:"foreignKey: DateID" json:"date"`
	// Time Time `gorm:"foreignKey: TimeID" json:"time"`
}

type SendEvent struct {
	GuestEmail string `json:"guestEmail"`
	HostEmail  string `json:"hostEmail"`
	Name       string `json:"name"`
	Date       string `json:"date"`
	Start      string `json:"start"`
	End        string `json:"end"`
	Room       string `json:"room"`
}

type GetEventPayload struct {
	AppointmentID uint `param:"appointmentID"`
}

type EventWithRelation struct {
	Event
	SendEvent
}

type SendEventPayload struct {
	AppointmentID uint `param:"appointmentID"`
}
