package middlewares

import (
	"c-drama-hub/config" // เรียกใช้ตรายางจากแผนกช่าง
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// AuthMiddleware ทำหน้าที่ตรวจบัตรก่อนเข้าโซน VIP
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "ไม่มีบัตร ห้ามเข้าครับ! ❌"})
			c.Abort()
			return
		}

		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		}

		// ใช้ config.JwtSecret เป็นตรายางในการสแกน
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return config.JwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "บัตรปลอม หรือบัตรหมดอายุครับ ❌"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("user_id", claims["user_id"])
			c.Set("role", claims["role"])
		}

		c.Next()
	}
}
