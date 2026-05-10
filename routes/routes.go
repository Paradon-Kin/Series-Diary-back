package routes

import (
	"series-diary/controllers"
	"series-diary/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// 🌍 โซนทั่วไป (Public)
	r.GET("/ping", func(c *gin.Context) { c.JSON(200, gin.H{"message": "API is running!"}) })
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	// 🌟 โซน VIP (Protected)
	api := r.Group("/api")
	api.Use(middlewares.AuthMiddleware())
	{
		api.POST("/series", controllers.AddSeries)
		api.GET("/series", controllers.GetSeries)

		api.POST("/history", controllers.UpdateHistory)
		api.GET("/history", controllers.GetHistory)

		api.DELETE("/series/:id", controllers.DeleteSeries)

		api.POST("/rate", controllers.UpdateRating)
	}

	return r
}
