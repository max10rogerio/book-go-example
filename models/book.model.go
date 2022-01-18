package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID          uint    `json:"id" gorm:"primaryKey"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	MediumPrice float32 `json:"medium_price" gorm:"type:decimal(10,2);"`
	Author      string  `json:"author"`
	ImageUrl    string  `json:"image_url"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}
