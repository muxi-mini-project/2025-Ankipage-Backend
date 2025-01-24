package controllers

import (
	"Ankipage/db"
	"Ankipage/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

var jwtKey = []byte("your_secret_key") // 替换为安全的密钥

// Token 数据结构
type Claims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

// @Summary Register a new user
// @Description Create a new user account with a username and password
// @Tags Authentication
// @Accept json
// @Produce json
// @Param user body object{username=string,password=string} true "User credentials"
// @Success 201 {object} map[string]interface{} "User registered successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request payload"
// @Failure 500 {object} map[string]interface{} "Failed to hash password or create user"
// @Router /register [post]
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
	c.JSON(http.StatusOK, user)
}

// 登录用户
// @Summary Login a user
// @Description Authenticate a user with username and password, return a JWT token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param user body object{username=string,password=string} true "User credentials"
// @Success 200 {object} map[string]interface{} "Logged in successfully, token generated"
// @Failure 400 {object} map[string]interface{} "Invalid request payload"
// @Failure 401 {object} map[string]interface{} "Invalid username or password"
// @Failure 500 {object} map[string]interface{} "Failed to generate token"
// @Router /login [post]
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
	c.JSON(http.StatusOK, user)
}
