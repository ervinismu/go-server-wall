package models

import (
	"github.com/jinzhu/gorm"
)

// Struct for generate table payment
type Payment struct {
	gorm.Model
	PaymentStatusID uint            `form:"paymentStatusId" json:"paymentStatusId"`
	PaymentStatus   []PaymentStatus `gorm:"ForeignKey:PaymentStatusID" form:"paymentStatus" json:"paymentStatus"`
	UserID          uint            `form:"userId" json:"userId"`
	WallID          uint            `form:"wallId" json:"wallId"`
}
