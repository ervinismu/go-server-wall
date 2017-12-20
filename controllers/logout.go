package controllers

import (
	"github.com/ervinismu/go-server-wall/template"
	"github.com/ervinismu/go-server-wall/models"
	"github.com/gin-gonic/gin"
)

// Function for Logout from apps
func Logout(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	var user models.User
	var res models.Response
	// get token from header
	token := c.Request.Header.Get("token")
	// check user request without token
	if token == "" {
		res.Code = "401"
		res.Message = "You are not login!"
		res.Status = "FAILED"
		data := template.Response(&res)
		c.JSON(200, data)
	} else if err := db.Select("token").Where("token = ?", token).First(&user).Error; err != nil {
		// check if token not in db
		res.Code = "401"
		res.Message = "Your token invalid!"
		res.Status = "FAILED"
		data := template.Response(&res)
		c.JSON(200, data)
	} else {
		// if contain token and token in db
		db.Model(&user).Select("token").Update(map[string]interface{}{"token": ""})
		res.Code = "200"
		res.Message = "Logout Succes!"
		res.Status = "SUCCESS"
		data := template.Response(&res)
		c.JSON(200, data)
	}
}