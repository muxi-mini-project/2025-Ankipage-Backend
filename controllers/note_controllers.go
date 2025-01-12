package controllers

import (
	"Ankipage/db"
	"Ankipage/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type CrtNote struct {
	title   string `json:"title" binding:"required"`
	content string `json:"content" binding:"required"`
}

// @Summary Create a new note
// @Description Create a note for a user
// @Tags Notes
// @Accept json
// @Produce json
// @Param userid path int true "User ID"
// @Param note body models.Note true "Note Data"
// @Success 200 {object} map[string]interface{} "Note created successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request body"
// @Failure 500 {object} map[string]interface{} "Failed to save note"
// @Router /notes/{userid} [post]
func CreateNote(c *gin.Context) {
	userID := c.Param("userid")
	UserID, _ := strconv.Atoi(userID)
	var note models.Note
	if err := c.BindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error: Invalid request body   ": err.Error()})
		return
	}
	note.ID = c.GetInt("user_id")
	note.UserID = UserID
	note.CreatedAt = time.Now()
	note.UpdatedAt = time.Now()
	// 保存到数据库
	if err := note.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save note"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Note created successfully", "note_id": note.ID})
}

// @Summary List recent notes
// @Description Get the most recent notes for a user
// @Tags Notes
// @Accept json
// @Produce json
// @Param userid path int true "User ID"
// @Success 200 {array} models.Note "List of recent notes"
// @Failure 500 {object} map[string]interface{} "Failed to fetch notes"
// @Router /notes/recent/{userid} [get]
func ListRecentNotes(c *gin.Context) {
	userID := c.GetInt("userid")
	limit := 4

	// 获取最近的笔记
	notes, err := models.GetRecentNotes(userID, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch notes"})
		return
	}

	c.JSON(http.StatusOK, notes)
}

// @Summary Get a specific note
// @Description Get a note by its ID
// @Tags Notes
// @Accept json
// @Produce json
// @Param id path int true "Note ID"
// @Success 200 {object} models.Note "Note data"
// @Failure 404 {object} map[string]interface{} "Note not found"
// @Router /notes/{id} [get]
func GetNote(c *gin.Context) {
	id := c.Param("id")
	noteID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid note id"})
	}
	// 获取笔记
	note, err := models.GetNoteByID(noteID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
		return
	}

	c.JSON(http.StatusOK, note)
}

// @Summary Get all notes
// @Description Get all notes for a user
// @Tags Notes
// @Accept json
// @Produce json
// @Param userid path int true "User ID"
// @Success 200 {array} models.Note "List of all notes"
// @Failure 404 {object} map[string]interface{} "Notes not found"
// @Router /notes/all/{userid} [get]
func GetNotes(c *gin.Context) {
	userID := c.Param("userid")
	UserID, _ := strconv.Atoi(userID)
	fmt.Println(UserID)
	// 获取笔记
	note, err := models.GetAllNotes(UserID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
		return
	}

	c.JSON(http.StatusOK, note)
}

// @Summary Update a note
// @Description Update a specific note by its ID
// @Tags Notes
// @Accept json
// @Produce json
// @Param id path int true "Note ID"
// @Param note body models.Note true "Updated Note Data"
// @Success 200 {object} map[string]interface{} "Note updated successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request body"
// @Failure 404 {object} map[string]interface{} "Note not found"
// @Failure 500 {object} map[string]interface{} "Failed to update note"
// @Router /notes/{id} [put]
func UpdateNote(c *gin.Context) {
	id := c.Param("id")
	var note models.Note

	// 查找原始笔记
	if err := db.DB.First(&note, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
		return
	}

	// 更新笔记内容
	if err := c.BindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	note.UpdatedAt = time.Now()

	if err := note.Update(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update note"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Note updated successfully"})
}

// @Summary Delete a note
// @Description Delete a specific note by its ID
// @Tags Notes
// @Accept json
// @Produce json
// @Param id path int true "Note ID"
// @Success 200 {object} map[string]interface{} "Note deleted successfully"
// @Failure 400 {object} map[string]interface{} "Invalid ID"
// @Failure 500 {object} map[string]interface{} "Failed to delete note"
// @Router /notes/{id} [delete]
func DeleteNote(c *gin.Context) {
	id := c.Param("id")
	noteID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	// 删除笔记
	if err := models.DeleteNoteByID(noteID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete note"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Note deleted successfully"})
}
