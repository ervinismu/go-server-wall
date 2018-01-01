package models

import (
	"github.com/jinzhu/gorm"
)

// Struct for generate table wall
type Wall struct {
	gorm.Model
	Payment     Payment `gorm:"ForeignKey:WallID" form:"payment" json:"payment"`
	Title       string  `form:"title" json:"title"`
	Description string  `form:"description" json:"description"`
	Price       string  `form:"price" json:"price"`
	Address     string  `form:"address" json:"address"`
}
