package model

import (
	"time"

	"gorm.io/gorm"
)

type Time struct {
	gorm.Model
	DateId uint      `gorm:"not null" json:"date_id"`
	Start  time.Time `gorm:"not null" json:"start"`
	End    time.Time `gorm:"not null" json:"end"`
}
