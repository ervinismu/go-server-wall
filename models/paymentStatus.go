package models

import (
	"github.com/jinzhu/gorm"
)

// Struct for generate payment status
type PaymentStatus struct {
	gorm.Model
	Payment     Payment `gorm:"ForeignKey:PaymentStatusID" form:"payment" json:"payment"`
	UserID      uint    `form:"userId" json:"userId"`
	Name        string  `form:"name" json:"name"`
	Description string  `form:"description" json:"description"`
}
