package tests

import (
	"bytes"
	"context"
	"devops_project/api"
	"devops_project/db/models"
	"devops_project/repository"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"

	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var h *api.Handler

func setupTestDB() (func(), error) {
	ctx := context.Background()

	// Create PostgreSQL container
	req := testcontainers.ContainerRequest{
		Image:        "postgres:latest",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_PASSWORD": "password",
			"POSTGRES_DB":       "testdb",
		},
		WaitingFor: wait.ForListeningPort("5432/tcp").WithStartupTimeout(60 * time.Second),
	}

	postgresC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		// panic(err)
		return nil, err
	}

	// Get the container's host and port
	host, _ := postgresC.Host(ctx)
	port, _ := postgresC.MappedPort(ctx, "5432")

	dsn := fmt.Sprintf("host=%s port=%s user=postgres password=password dbname=testdb sslmode=disable", host, port.Port())

	// Connect to the PostgreSQL database
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		// panic("failed to connect to database")
		return nil, err
	}

	// Migrate the schema
	db.AutoMigrate(&models.User{})

	// Initialize the handler
	repo := repository.NewRepo(db)
	h = api.NewHandler(repo)

	// Clean up database before each test
	db.Exec("DELETE FROM users")

	// Tear down the container after tests
	// defer postgresC.Terminate(ctx)
	cleanup := func() {
		postgresC.Terminate(ctx)
	}

	return cleanup, nil
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	api.RegisterRoutes(r, db)
	return r
}

func TestMain(m *testing.M) {
	cleanup, err := setupTestDB()
	if err != nil {
		log.Fatalf("failed to setup test DB: %v", err)
	}
	code := m.Run()
	cleanup()
	os.Exit(code)
}

func TestCreateUserHandler(t *testing.T) {
	r := setupRouter()

	payload := models.UserInput{
		Name:  "Bob",
		Email: "bob@example.com",
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "User created successfully")
}

func TestGetAllUsersHandler(t *testing.T) {
	r := setupRouter()

	// Create a user directly in DB for testing
	_ = db.Exec("INSERT INTO users (name, email) VALUES (?, ?)", "Alice", "alice@example.com")

	req, _ := http.NewRequest("GET", "/users", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "alice@example.com")
}

func TestGetUserHandler(t *testing.T) {
	r := setupRouter()

	_ = db.Exec("INSERT INTO users (name, email) VALUES (?, ?)", "Carl", "carl@example.com")

	var user models.User
	db.Last(&user)

	req, _ := http.NewRequest("GET", "/users/"+fmt.Sprint(user.Id), nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "carl@example.com")
}

func TestUpdateUserHandler(t *testing.T) {
	r := setupRouter()

	_ = db.Exec("INSERT INTO users (name, email) VALUES (?, ?)", "Dave", "dave@example.com")

	var user models.User
	db.Last(&user)

	payload := models.UserInput{Name: "David", Email: "david@example.com"}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("PUT", "/users/"+fmt.Sprint(user.Id), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "David")
}

func TestDeleteUserHandler(t *testing.T) {
	r := setupRouter()

	_ = db.Exec("INSERT INTO users (name, email) VALUES (?, ?)", "Eve", "eve@example.com")

	var user models.User
	db.Last(&user)

	req, _ := http.NewRequest("DELETE", "/users/"+fmt.Sprint(user.Id), nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}
