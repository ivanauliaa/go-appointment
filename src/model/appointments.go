package model

import "gorm.io/gorm"

type Appointment struct {
	gorm.Model
	HostID  uint  `gorm:"not null" json:"host_id"`
	GuestID uint  `json:"guest_id"`
	URL     URL   `gorm:"foreignKey:AppointmentID" json:"url"`
	Dates   Date  `gorm:"foreignKey:AppointmentID" json:"dates"`
	Host    User  `gorm:"foreignKey:HostID" json:"host"`
	Guest   User  `gorm:"foreignKey:GuestID" json:"guest"`
	Event   Event `gorm:"foreignKey:AppointmentID"`
}
