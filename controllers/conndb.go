package controllers

import (
	"fmt"
	"os"

	"github.com/ervinismu/go-server-wall/models"
	"github.com/jinzhu/gorm"
	// for connect postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Function for connected db
func InitDb() *gorm.DB {

	var err error
	var user models.User
	var model models.Model
	var paymentStatus models.PaymentStatus
	var userDetail models.UserDetail
	var wall models.Wall

	host := os.Getenv("DATABASE_URL")
	if host == "" {
		host = "host=127.0.0.1 user=postgres dbname=walldb sslmode=disable password=walldb"
	}

	db, err := gorm.Open("postgres", host)
	db.AutoMigrate(user, model, paymentStatus, userDetail, wall)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Db connected!")
	}
	return db
}
