package main

import (
  "github.com/ervinismu/go-server-wall/controllers"
	"github.com/gin-gonic/gin"
	"os"
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
  
  r.Run(":" + port)
}
