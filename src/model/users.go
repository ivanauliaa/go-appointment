package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"not null" json:"name" validate:"required"`
	Email    string `gorm:"not null;unique" json:"email" validate:"required,email"`
	Password string `gorm:"not null" json:"password" validate:"required"`
	// Appointments []Appointment `gorm:"foreignKey:HostID" json:"appointments"`
}
