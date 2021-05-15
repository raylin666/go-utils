package database

import "time"

type Model struct {
	ID        uint64 `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
