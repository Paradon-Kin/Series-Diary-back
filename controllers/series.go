package controllers

import (
	"net/http"
	"series-diary/config"
	"series-diary/models"

	"github.com/gin-gonic/gin"
)

func AddSeries(c *gin.Context) {
	var input models.Series
	c.ShouldBindJSON(&input)

	// 1. ดึง ID ของคนที่ล็อกอินอยู่
	userIDFloat, _ := c.Get("user_id")
	input.UserID = uint(userIDFloat.(float64)) // 👈 ประทับตราว่าใครเป็นคนเพิ่มซีรีส์เรื่องนี้

	config.DB.Create(&input)
	c.JSON(http.StatusOK, gin.H{"message": "เพิ่มซีรีส์สำเร็จ", "data": input})
}

func GetSeries(c *gin.Context) {
	// 1. ดึง ID ของคนที่ล็อกอินอยู่
	userIDFloat, _ := c.Get("user_id")
	userID := uint(userIDFloat.(float64))

	var seriesList []models.Series
	// 2. 🚨 ค้นหาเฉพาะซีรีส์ที่เป็นของ User คนนี้เท่านั้น
	config.DB.Where("user_id = ?", userID).Find(&seriesList)

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

func UpdateRating(c *gin.Context) {
	// 1. รับค่า SeriesID และ คะแนนที่ต้องการให้
	var input struct {
		SeriesID  uint    `json:"series_id"`
		AvgRating float64 `json:"avg_rating"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ข้อมูลไม่ถูกต้อง"})
		return
	}

	// 2. ตรวจสอบว่าใครกำลังล็อกอินอยู่
	userIDFloat, _ := c.Get("user_id")
	userID := uint(userIDFloat.(float64))

	var series models.Series
	// 3. ค้นหาซีรีส์เรื่องนั้น โดยต้องแน่ใจว่าเป็นของ User คนนี้จริงๆ
	if err := config.DB.Where("id = ? AND user_id = ?", input.SeriesID, userID).First(&series).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ไม่พบซีรีส์ หรือคุณไม่มีสิทธิ์แก้ไข"})
		return
	}

	// 4. อัปเดตคะแนนลงฟิลด์ AvgRating โดยตรง
	series.AvgRating = input.AvgRating
	config.DB.Save(&series)

	c.JSON(http.StatusOK, gin.H{
		"message": "อัปเดตเรตติ้งสำเร็จ",
		"data":    series,
	})
}

func DeleteSeries(c *gin.Context) {
	// 1. รับ ID จาก URL Parameter (เช่น /api/series/5)
	id := c.Param("id")

	// 2. ตรวจสอบว่าใครกำลังล็อกอินอยู่
	userIDFloat, _ := c.Get("user_id")
	userID := uint(userIDFloat.(float64))

	var series models.Series
	// 3. ค้นหาซีรีส์เรื่องนั้น โดยต้องเป็นของ User คนนี้เท่านั้นถึงจะอนุญาต
	if err := config.DB.Where("id = ? AND user_id = ?", id, userID).First(&series).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ไม่พบซีรีส์ที่ต้องการลบ หรือคุณไม่มีสิทธิ์ลบรายการนี้"})
		return
	}

	// 4. ลบข้อมูล (GORM จะทำ Soft Delete ให้ตามที่เราใช้ gorm.Model)
	config.DB.Delete(&series)

	c.JSON(http.StatusOK, gin.H{"message": "ลบซีรีส์ออกจากไดอารี่เรียบร้อยแล้ว"})
}
