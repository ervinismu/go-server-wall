package main

import (
	"github.com/gin-gonic/gin"
	"os"
)

func main() {  
  port := os.Getenv("PORT")
  if port == "" {
    port = "8080"
  }

  r := gin.Default()
  r.GET("/", func(c *gin.Context) {
    c.String(200, "Hello World bro! this is first deploy automation heroku broo!")
  })
  r.Run(":" + port)
}
