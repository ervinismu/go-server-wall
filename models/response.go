package models

// Struct for generate Response format
type Response struct {
	Code	string	`form:"code" json:"code"`
	Message	string	`form:"message" json:"message"`
	Token	string	`form:"token" json:"token"`
	Status	string	`form:"status" json:"status"`
}