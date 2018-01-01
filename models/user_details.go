package models

import (
	"github.com/jinzhu/gorm"
)

// Struct for generate table user detail
type UserDetail struct {
	gorm.Model
	UserID  uint   `form:"userId" json:"userId"`
	Name    string `form:"name" json:"name"`
	Address string `form:"address" json:"address"`
	Company string `form:"company" json:"company"`
	Gender  string `form:"gender" json:"gender"`
}
