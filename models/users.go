package models

import (
	"github.com/jinzhu/gorm"
)

// Struct for createUser
type User struct {
	gorm.Model
	UserDetail      UserDetail `gorm:"ForeignKey:UserID" form:"userDetail" json:"userDetail"`
	Roles           Roles      `gorm:"ForeignKey:UserID" form:"roles" json:"roles"`
	Payment         []Payment  `gorm:"ForeignKey:UserID" form:"payment" json:"payment"`
	Name            string     `form:"name" json:"name"`
	Email           string     `form:"email" json:"email"`
	Password        string     `form:"password" json:"password"`
	PasswordConfirm string     `form:"passwordConfirm" json:"passwordConfirm"`
	Token           string     `form:"token" json:"token"`
}
