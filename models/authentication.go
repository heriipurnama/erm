package models

import "time"

type Authentication struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"unique;size:255"`
	Password  string `gorm:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
