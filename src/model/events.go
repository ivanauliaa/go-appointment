package model

import "gorm.io/gorm"

type Event struct {
	gorm.Model
	AppointmentID uint   `gorm:"not null" json:"appointment_id"`
	DateID        uint   `json:"date_id"`
	TimeID        uint   `json:"time_id"`
	Status        string `gorm:"not null" json:"status"`
	// Appointment   Appointment `gorm:"foreignKey: AppointmentID" json:"appointment"`
	Date Date `gorm:"foreignKey: DateID" json:"date"`
	Time Time `gorm:"foreignKey: TimeID" json:"time"`
}
