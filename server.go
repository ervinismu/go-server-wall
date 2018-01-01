package main

import (
	"os"

	"github.com/ervinismu/go-server-wall/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World bro! this is first deploy automation heroku broo!")
	})
	r.POST("/login", controllers.Login)
	r.POST("/register", controllers.Register)
	r.GET("/logout", controllers.Logout)

	r.GET("/wall", controllers.GetWall)
	r.GET("/wall/:id", controllers.GetDetailWall)
	r.DELETE("/wall/:id", controllers.DeleteWall)
	r.PUT("/wall/:id", controllers.UpdateWall)

	r.GET("/user", controllers.GetAllUser)
	r.GET("/user/:id", controllers.GetDetailUser)
	r.DELETE("/user/:id", controllers.DeleteUser)
	r.PUT("/user/:id", controllers.UpdateUser)

	r.Run(":" + port)
}
