// package main

// import (
// 	"net/http"
// 	"strconv"
// 	"sync"

// 	"github.com/gin-gonic/gin"
// )

// type User struct {
// 	ID    int    `json:"id"`
// 	Name  string `json:"name"`
// 	Email string `json:"email"`
// }

// var (
// 	users   = make(map[int]User)
// 	nextID  = 1
// 	userMux sync.Mutex
// )

// func main() {
// 	r := gin.Default()

// 	// Welcome route
// 	r.GET("/", func(c *gin.Context) {
// 		c.String(http.StatusOK, "Welcome to your Gin API!")
// 	})

// 	// Create user
// 	r.POST("/users", func(c *gin.Context) {
// 		var input struct {
// 			Name  string `json:"name"`
// 			Email string `json:"email"`
// 		}

// 		if err := c.ShouldBindJSON(&input); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		userMux.Lock()
// 		newUser := User{
// 			ID:    nextID,
// 			Name:  input.Name,
// 			Email: input.Email,
// 		}
// 		users[nextID] = newUser
// 		nextID++
// 		userMux.Unlock()

// 		c.JSON(http.StatusCreated, newUser)
// 	})

// 	// Read all users
// 	r.GET("/users", func(c *gin.Context) {
// 		userMux.Lock()
// 		defer userMux.Unlock()

// 		result := make([]User, 0, len(users))
// 		for _, user := range users {
// 			result = append(result, user)
// 		}
// 		c.JSON(http.StatusOK, result)
// 	})

// 	// Read single user
// 	r.GET("/users/:id", func(c *gin.Context) {
// 		id, err := strconv.Atoi(c.Param("id"))
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
// 			return
// 		}

// 		userMux.Lock()
// 		user, exists := users[id]
// 		userMux.Unlock()

// 		if !exists {
// 			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
// 			return
// 		}
// 		c.JSON(http.StatusOK, user)
// 	})

// 	// Update user
// 	r.PUT("/users/:id", func(c *gin.Context) {
// 		id, err := strconv.Atoi(c.Param("id"))
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
// 			return
// 		}

// 		var input struct {
// 			Name  string `json:"name"`
// 			Email string `json:"email"`
// 		}

// 		if err := c.ShouldBindJSON(&input); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		userMux.Lock()
// 		defer userMux.Unlock()

// 		if _, exists := users[id]; !exists {
// 			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
// 			return
// 		}

// 		updatedUser := User{
// 			ID:    id,
// 			Name:  input.Name,
// 			Email: input.Email,
// 		}
// 		users[id] = updatedUser

// 		c.JSON(http.StatusOK, updatedUser)
// 	})

// 	// Delete user
// 	r.DELETE("/users/:id", func(c *gin.Context) {
// 		id, err := strconv.Atoi(c.Param("id"))
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
// 			return
// 		}

// 		userMux.Lock()
// 		defer userMux.Unlock()

// 		if _, exists := users[id]; !exists {
// 			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
// 			return
// 		}

// 		delete(users, id)
// 		c.Status(http.StatusNoContent)
// 	})

//		r.Run(":8080")
//	}
package main

import (
	"devops_project/api"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	api.RegisterRoutes(r)
	r.Run(":8080")
}
