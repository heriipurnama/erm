package models

import "time"

type Order struct {
	ID         uint `gorm:"primaryKey"`
	CustomerID uint
	Product    string `gorm:"size:255"`
	Quantity   int
	Price      float64
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
