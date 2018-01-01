package models

import (
	"time"
)

// Struct for generate Model
type Model struct {
	ID        uint `form:"id" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
