package database

type Model struct {
	ID        uint64 `gorm:"primary_key" json:"id"`
	CreatedAt uint64 `json:"created_at"`
	UpdatedAt uint64 `json:"updated_at"`
}
