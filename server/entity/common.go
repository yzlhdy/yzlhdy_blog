package entity

import "time"

type Model struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *string   `gorm:"default:null;index:dbc" json:"deleted_at"`
}
