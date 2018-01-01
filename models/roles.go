package models

import (
	"github.com/jinzhu/gorm"
)

// Struct for generate table roles
type Roles struct {
	gorm.Model
	UserID      uint   `form:"userId" json:"userId"`
	Name        string `form:"name" json:"name"`
	Description string `form:"description" json:"description"`
}
