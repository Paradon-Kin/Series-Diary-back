package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
	Role     string `gorm:"default:'user'"`
}

type Series struct {
	gorm.Model
	Title         string
	Genre         string
	TotalEpisodes int
	CoverURL      string
	AvgRating     float64 `gorm:"default:0"`
	UserID        uint
}

type WatchHistory struct {
	gorm.Model
	UserID         uint   `json:"user_id"`
	SeriesID       uint   `json:"series_id"`
	CurrentEpisode int    `json:"current_episode"`
	Series         Series `gorm:"foreignKey:SeriesID" json:"series"`
}
