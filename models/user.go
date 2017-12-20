package models

// Struct for createUser
type User struct {
	ID				uint 	`form:"id" json:"id"`
	Name			string	`form:"name" json:"name"`
	Email			string	`form:"email" json:"email"`
	Password		string	`form:"password" json:"password"`
	PasswordConfirm	string	`form:"passwordConfirm" json:"passwordConfirm`
	Token			string	`form:"token" json:"token"`
}