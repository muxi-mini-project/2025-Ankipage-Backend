package controllers

import (
	"Ankipage/dataoperation"
	"Ankipage/db"
	"Ankipage/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

func RegisterUser(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// 创建用户
	user := models.User{
		Username: input.Username,
		Password: string(hashedPassword),
	}
	if err := db.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}
func LoginUser(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := db.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// 生成 JWT Token
	expirationTime := time.Now().Add(24 * time.Hour) // 设置过期时间
	claims := &Claims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	// 设置 Cookie
	c.SetCookie("token", tokenString, int(expirationTime.Sub(time.Now()).Seconds()), "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "Logged in successfully"})
}

type CrtNote struct {
	title   string `json:"title" binding:"required"`
	content string `json:"content" binding:"required"`
}

func CreateNote(c *gin.Context) {
	var note models.Note
	if err := c.BindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error: Invalid request body   ": err.Error()})
		return
	}
	note.CreatedAt = time.Now()
	note.UpdatedAt = time.Now()

	// 保存到存储系统 (示例使用内存)
	dataoperation.SaveNote(&note)
	c.JSON(http.StatusOK, gin.H{"message": "Note created successfully", "note_id": note.ID})
}
func ListRecentNotes(c *gin.Context) {

}
func GetNotes(c *gin.Context) {
	
}
func UpdateNote(c *gin.Context) {

}
func DeleteNote(c *gin.Context) {

}
