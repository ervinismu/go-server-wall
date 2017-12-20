package controllers

import (
	"crypto/sha256"
	"encoding/base64"
	"github.com/ervinismu/go-server-wall/models"
	"github.com/ervinismu/go-server-wall/template"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Function for Login
func Login(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	var user models.User
	// get email dan password
	c.Bind(&user)
	password := user.Password
	// check email
	if err := db.Select("email, password").Where("email = ?", user.Email).First(&user).Error; err != nil {
		// if err, skip generate token and return response error
		var res models.Response
		res.Code  =  "401"
		res.Message  =  "Email doesn't exist!"
		res.Status = "FAILED"
		data := template.Response(&res)
		c.JSON(401, data)
	} else {
		if CheckPasswordHash(password, user.Password) {
			// Generate token
			originalText := user.Email + user.Password
			user.Token = GenerateToken(originalText)
			// Save token into db
			db.Model(&user).Where("email = ?", user.Email).Update("token", user.Token)
			// Return token and message
			var res models.Response
			res.Code = "201"
			res.Message = "Login Success!"
			res.Status = "SUCCESS"
			res.Token = user.Token
			data := template.Response(&res)
			c.JSON(200, data)
			return
		}
		var res models.Response
		res.Code = "401"
		res.Message = "Your password is wrong!"
		res.Token = user.Token
		res.Status = "FAILED"
		data := template.Response(&res)
		c.JSON(400, data)
	}
}

// CheckPasswordHash is using check password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Function for generate Token
func GenerateToken(text string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(text))
	h := sha256.New()
	h.Write([]byte(encoded))
	var sha =  h.Sum(nil)
	var token = base64.StdEncoding.EncodeToString([]byte(sha))
	return token
}