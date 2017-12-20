package controllers

import (
	"golang.org/x/crypto/bcrypt"
	"regexp"
	"github.com/ervinismu/go-server-wall/template"
	"github.com/ervinismu/go-server-wall/models"
	"github.com/gin-gonic/gin"
)

// Function for Register new user
func Register(c *gin.Context) {
	checkRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	db := InitDb()
	defer db.Close()
	var user models.User
	var res models.Response
	c.Bind(&user)
	// create validation register
	if user.Name == "" {
		res.Code = "401"
		res.Message = "Field can not be empty!"
		res.Status = "FAILED"
		data := template.Response(&res)
		c.JSON(400, data)
	} else if user.Password == "" {
		res.Code = "401"
		res.Message = "Field can not be empty!"
		res.Status = "FAILED"
		data := template.Response(&res)
		c.JSON(400, data)
	} else if  user.Password != user.PasswordConfirm {
		res.Code = "401"
		res.Message = "Password confirmation not same!"
		res.Status = "FAILED"
		data := template.Response(&res)
		c.JSON(400, data)
	} else if !checkRegex.MatchString(user.Email) {
		res.Code = "401"
		res.Message = "Invalid format email!"
		res.Status = "FAILED"
		data := template.Response(&res)
		c.JSON(400, data)
	} else if err := db.Where("email = ?", user.Email).First(&user).Error; err != nil {
		plainPassword := user.Password
		hashPassword, err := HashPassword(user.Password)
		if err != nil {
			res.Code = "401"
			res.Message = "Failed encrypt password!"
			res.Status = "FAILED"
			data := template.Response(&res)
			c.JSON(400, data)
		}
		user.Password = string(hashPassword)
		user.PasswordConfirm = string("")
		// insert into db after encrypt
		db.Save(&user)
		if CheckPasswordHash(plainPassword, user.Password) {
			// senerate token
			originalText := user.Email + user.Password
			user.Token = GenerateToken(originalText)
			// save token into db
			db.Model(&user).Where("email = ?", user.Email).Update("token", user.Token)
			// return response with token
			var res models.Response
			res.Code = "201"
			res.Message = "Register Success!"
			res.Status = "SUCCESS"
			res.Token = user.Token
			data := template.Response(&res)
			c.JSON(200, data)
			return
		}
	} else {
		res.Code = "401"
		res.Message = "Email has been used!"
		res.Status = "FAILED"
		data := template.Response(&res)
		c.JSON(400, data)
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}