package controllers

import (
	"c-drama-hub/config"
	"c-drama-hub/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddSeries(c *gin.Context) {
	var input models.Series
	c.ShouldBindJSON(&input)
	config.DB.Create(&input)
	c.JSON(http.StatusOK, gin.H{"message": "เพิ่มซีรีส์สำเร็จ", "data": input})
}

func GetSeries(c *gin.Context) {
	var seriesList []models.Series
	config.DB.Find(&seriesList)
	c.JSON(http.StatusOK, gin.H{"data": seriesList})
}

func UpdateHistory(c *gin.Context) {
	var input struct {
		SeriesID       uint `json:"series_id"`
		CurrentEpisode int  `json:"current_episode"`
	}
	c.ShouldBindJSON(&input)

	userIDFloat, _ := c.Get("user_id")
	userID := uint(userIDFloat.(float64))

	var history models.WatchHistory
	if config.DB.Where("user_id = ? AND series_id = ?", userID, input.SeriesID).First(&history).Error != nil {
		history = models.WatchHistory{UserID: userID, SeriesID: input.SeriesID, CurrentEpisode: input.CurrentEpisode}
		config.DB.Create(&history)
	} else {
		history.CurrentEpisode = input.CurrentEpisode
		config.DB.Save(&history)
	}
	c.JSON(http.StatusOK, gin.H{"message": "อัปเดตประวัติสำเร็จ", "data": history})
}

func GetHistory(c *gin.Context) {
	userIDFloat, _ := c.Get("user_id")
	userID := uint(userIDFloat.(float64))

	var histories []models.WatchHistory
	config.DB.Preload("Series").Where("user_id = ?", userID).Find(&histories)
	c.JSON(http.StatusOK, gin.H{"data": histories})
}
