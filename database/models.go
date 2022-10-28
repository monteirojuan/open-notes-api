package database

import "time"

type Note struct {
	ID          uint   `gorm:"primaryKey"`
	title       string `gorm:"not null"`
	content     string
	pinned      bool      `gorm:"default:false"`
	archived    bool      `gorm:"default:false"`
	created_at  time.Time `gorm:"autoCreateTime"`
	updated_at  time.Time `gorm:"autoUpdateTime"`
	pinned_at   time.Time
	archived_at time.Time
}
