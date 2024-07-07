package models

import "time"

type Customer struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:255"`
	Email     string `gorm:"unique;size:255"`
	Address   string `gorm:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
