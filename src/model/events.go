package model

import "gorm.io/gorm"

type Event struct {
	gorm.Model
	AppointmentID uint `gorm:"not null" json:"appointmentID"`
	HostID        uint `gorm:"not null" json:"hostID"`
	GuestID       uint `gorm:"not null" json:"guestID"`
	// Name          string `gorm:"not null" json:"name"`
	DateID uint   `json:"dateID"`
	TimeID uint   `json:"timeID"`
	Status string `gorm:"not null" json:"status"`
	// Appointment   Appointment `gorm:"foreignKey: AppointmentID" json:"appointment"`
	// Date Date `gorm:"foreignKey: DateID" json:"date"`
	// Time Time `gorm:"foreignKey: TimeID" json:"time"`
}
