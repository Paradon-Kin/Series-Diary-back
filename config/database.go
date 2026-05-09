package config

import (
	"log"
	"os" // แพ็กเกจสำหรับอ่านไฟล์ .env

	"series-diary/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var JwtSecret []byte // ปรับให้กลายเป็นตัวแปรเปล่าๆ รอรับค่าจาก .env

func ConnectDatabase() {
	// 1. ดึงความลับออกจากตู้เซฟ (os.Getenv)
	dsn := os.Getenv("DB_DSN")
	JwtSecret = []byte(os.Getenv("JWT_SECRET"))

	// 2. เอาความลับไปไขกุญแจ
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ เชื่อมต่อ Database ไม่สำเร็จ: ", err)
	}

	log.Println("✅ เชื่อมต่อ Database สำเร็จแล้ว!")

	database.AutoMigrate(&models.User{}, &models.Series{}, &models.WatchHistory{})
	log.Println("✅ สร้างตารางเรียบร้อย!")

	DB = database
}
