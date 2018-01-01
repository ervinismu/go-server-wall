package controllers

import (
	"fmt"

	"github.com/ervinismu/go-server-wall/models"
	"github.com/gin-gonic/gin"
)

// Get all wall data
func GetWall(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	var wall []models.Wall
	if err := db.Raw("SELECT * FROM walls ORDER BY id DESC").Scan(&wall).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	data := map[string]interface{}{
		"data": wall,
	}
	c.JSON(200, data)
}

// Get wall detail
func GetDetailWall(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	id := c.Params.ByName("id")
	var wall models.Wall
	if err := db.Where("id = ?", id).First(&wall).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(404)
	}
	data := map[string]interface{}{
		"data": wall,
	}
	c.JSON(200, data)
}

// Delete wall
func DeleteWall(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	id := c.Params.ByName("id")
	var wall models.Wall
	var res models.Response
	if err := db.Delete(&wall).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(404)
	}
	res.Status = "OKE"
	res.Message = "Wall " + id + " has been deleted."
	res.Code = "201"
	data := map[string]interface{}{
		"data": res,
	}
	c.JSON(200, data)
}

// Update wall
func UpdateWall(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	id := c.Params.ByName("id")
	var wall models.Wall
	var res models.Response
	if err := db.Where("id = ?", id).First(&wall).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(404)
	}
	c.BindJSON(&wall)
	db.Save(&wall)
	res.Code = "201"
	res.Message = "Update Success!"
	res.Status = "OKE"
	data := map[string]interface{}{
		"data": res,
	}
	c.JSON(200, data)
}
