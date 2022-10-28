package database

import "time"

type Note struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"not null" json:"title"`
	Content   string    `json:"content"`
	Pinned    bool      `gorm:"default:false" json:"pinned"`
	Archived  bool      `gorm:"default:false" json:"archived"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
