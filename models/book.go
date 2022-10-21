package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model

	Title       string `gorm:"type:varchar(250);not null" json:"title"`
	Description string `gorm:"not null" json:"description"`
	Author      string `gorm:"type:varchar(250);not null" json:"author"`
	Publication string `gorm:"type:varchar(250);not null" json:"publication"`
}
