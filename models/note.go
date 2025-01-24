package models

import (
	"Ankipage/db"
	"time"
)

type Note struct {
	ID        int       `gorm:"primaryKey"`
	UserID    int       `gorm:"index"`
	Title     string    `gorm:"size:255;not null"`
	Content   string    `gorm:"type:text"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

// 保存笔记
func (n *Note) Save() error {
	return db.DB.Create(n).Error
}

// 更新笔记
func (n *Note) Update() error {
	return db.DB.Save(n).Error
}

// 删除笔记
func DeleteNoteByID(id int) error {
	return db.DB.Delete(&Note{}, id).Error
}

// 获取笔记详情
func GetNoteByID(id int) (*Note, error) {
	var note Note
	err := db.DB.First(&note, id).Error
	return &note, err
}
func GetAllNotes(userID int) ([]Note, error) {
	var notes []Note
	err := db.DB.Where("user_id = ?", userID).Order("created_at DESC").Find(&notes).Error
	return notes, err
}

// 获取最近的 n 条笔记
func GetRecentNotes(userID, limit int) ([]Note, error) {
	var notes []Note
	err := db.DB.Where("user_id = ?", userID).Order("created_at DESC").Limit(limit).Find(&notes).Error
	return notes, err
}
