package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome to your Gin API!")
	})

	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, world!")
	})

	r.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.String(200, "User ID: %s", id)
	})

	r.Run(":8080") // default is 0.0.0.0:8080
}
