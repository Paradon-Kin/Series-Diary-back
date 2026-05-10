package main

import (
	"log"
	"os"

	"series-diary/config"
	"series-diary/routes"

	"github.com/joho/godotenv"
)

func main() {
	// 0. โหลดตู้เซฟ (.env) เป็นอันดับแรกสุด!
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ ไม่พบไฟล์ .env หรือโหลดไม่สำเร็จ")
	}

	// 1. เปิดโกดัง
	config.ConnectDatabase()

	// 2. จัดแผนที่เส้นทาง
	r := routes.SetupRouter()

	// 3. ดึงพอร์ตจากตู้เซฟมาใช้ (เผื่ออนาคตอยากเปลี่ยนพอร์ต จะได้แก้ที่ .env ที่เดียว)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // เผื่อลืมตั้งค่าใน .env ให้ใช้ 8080 เป็นค่าสำรอง
	}

	// 4. เปิดร้าน!
	r.Run(":" + port)
}
