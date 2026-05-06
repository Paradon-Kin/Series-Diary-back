package routes

import (
	"c-drama-hub/controllers"
	"c-drama-hub/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 🌍 โซนทั่วไป (Public)
	r.GET("/ping", func(c *gin.Context) { c.JSON(200, gin.H{"message": "API is running!"}) })
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	// 🌟 โซน VIP (Protected)
	api := r.Group("/api")
	api.Use(middlewares.AuthMiddleware()) // เอายามมาเฝ้าที่ประตูด้านหน้านี้
	{
		api.POST("/series", controllers.AddSeries)
		api.GET("/series", controllers.GetSeries)

		api.POST("/history", controllers.UpdateHistory)
		api.GET("/history", controllers.GetHistory)
	}

	return r
}
