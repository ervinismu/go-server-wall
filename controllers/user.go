package controllers

import (
	"fmt"

	"github.com/ervinismu/go-server-wall/models"
	"github.com/gin-gonic/gin"
)

// Get all user
func GetAllUser(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	var user []models.User
	if err := db.Raw("SELECT * FROM users ORDER BY id DESC").Scan(&user).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(404)
	}
	data := map[string]interface{}{
		"data": user,
	}
	c.JSON(200, data)
}

// Get detail user
func GetDetailUser(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	var user models.User
	id := c.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(404)
	}
	data := map[string]interface{}{
		"data": user,
	}
	c.JSON(200, data)
}

// Delete user
func DeleteUser(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	id := c.Params.ByName("id")
	var user models.User
	var res models.Response
	db.Where("id = ?", id).First(&user)
	if err := db.Delete(&user).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(404)
	}
	res.Code = "201"
	res.Message = "User " + id + " has been deleted."
	res.Status = "OKE"
	data := map[string]interface{}{
		"data": res,
	}
	c.JSON(200, data)
}

// Create User
func CreateUser(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	var user models.User
	var res models.Response
	c.Bind(&user)
	if err := db.Create(&user).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(404)
	}
	res.Code = "201"
	res.Message = "Success inserted user!"
	res.Status = "OKE"
	data := map[string]interface{}{
		"data": res,
	}
	c.JSON(200, data)
}

//  Update user
func UpdateUser(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	var user models.User
	var res models.Response
	id := c.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(404)
	}
	c.BindJSON(&user)
	if err := db.Save(&user).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(404)
	}
	res.Code = "201"
	res.Message = "Success Update user."
	res.Status = "OKE"
	data := map[string]interface{}{
		"data": res,
	}
	c.JSON(200, data)
}
