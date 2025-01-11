package dataoperation

import (
	"Ankipage/models"
	"sort"
)

var notes = []models.Note{}
var idCounter = 1

// 保存笔记
func SaveNote(note *models.Note) {
	note.ID = uint(idCounter)
	idCounter++
	notes = append(notes, *note)
}

// 获取最近 n 条笔记
func GetRecentNotes(limit int) []models.Note {
	// 按更新时间倒序排序
	sort.Slice(notes, func(i, j int) bool {
		return notes[i].UpdatedAt.After(notes[j].UpdatedAt)
	})

	// 返回限制数量的笔记
	if len(notes) > limit {
		return notes[:limit]
	}
	return notes
}
