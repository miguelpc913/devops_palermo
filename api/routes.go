package api

import (
	"devops_project/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	repo := repository.NewRepo(db)
	handler := NewHandler(repo)

	r.GET("/", handler.welcome)
	r.GET("/panic", func(c *gin.Context) {
		panic("this is a panic for testing purposes")
	})
	r.POST("/users", handler.createUser)
	r.GET("/users", handler.getAllUsers)
	r.GET("/users/:id", handler.getUser)
	r.PUT("/users/:id", handler.updateUser)
	r.DELETE("/users/:id", handler.deleteUser)
}
