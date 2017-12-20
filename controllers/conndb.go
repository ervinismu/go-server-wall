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

	host := os.Getenv("DATABASE_URL")
	if host == "" {
		host = "host=127.0.0.1 user=postgres dbname=walldb sslmode=disable password=walldb"
	}

	db, err := gorm.Open("postgres", host)
	db.AutoMigrate(user)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Db connected!")
	}
	return db
}
