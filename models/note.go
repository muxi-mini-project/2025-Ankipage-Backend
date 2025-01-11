package models

import "time"

type Note struct {
	ID        uint      `gorm:"primaryKey"`
	Title     string    `gorm:"size:255;not null"`
	Content   string    `gorm:"type:text"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
