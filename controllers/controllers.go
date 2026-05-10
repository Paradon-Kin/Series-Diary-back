package controllers

import (
	"net/http"
	"time"

	"series-diary/config"
	"series-diary/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ข้อมูลไม่ถูกต้อง"})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 10)
	user := models.User{Username: input.Username, Password: string(hashedPassword)}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ชื่อผู้ใช้นี้มีคนใช้แล้ว!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "สมัครสมาชิกสำเร็จ🎉"})
}

func Login(c *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	c.ShouldBindJSON(&input)

	var user models.User
	if err := config.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "ไม่พบผู้ใช้นี้"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "รหัสผ่านผิด"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, _ := token.SignedString(config.JwtSecret)
	c.JSON(http.StatusOK, gin.H{"message": "เข้าสู่ระบบสำเร็จ!", "token": tokenString})
}
