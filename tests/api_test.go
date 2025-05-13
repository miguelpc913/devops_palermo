package Tests

import (
	"bytes"
	"devops_project/api"
	"devops_project/storage"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupTestRouter() *gin.Engine {
	// Reset in-memory storage before each test
	gin.SetMode(gin.ReleaseMode)
	storage.Reset()
	r := gin.Default()
	api.RegisterRoutes(r)
	return r
}

func TestCreateAndGetUser(t *testing.T) {
	router := setupTestRouter()

	// Create user
	body := map[string]string{"name": "Alice", "email": "alice@example.com"}
	jsonBody, _ := json.Marshal(body)

	req := httptest.NewRequest("POST", "/users", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusCreated, res.Code)

	var created map[string]interface{}
	json.Unmarshal(res.Body.Bytes(), &created)
	id := int(created["id"].(float64))

	// Get user
	getReq := httptest.NewRequest("GET", "/users/"+strconv.Itoa(id), nil)
	getRes := httptest.NewRecorder()
	router.ServeHTTP(getRes, getReq)

	assert.Equal(t, http.StatusOK, getRes.Code)
	var fetched map[string]interface{}
	json.Unmarshal(getRes.Body.Bytes(), &fetched)

	assert.Equal(t, "Alice", fetched["name"])
}

func TestGetAllUsers(t *testing.T) {
	router := setupTestRouter()

	// Create multiple users
	users := []map[string]string{
		{"name": "User1", "email": "u1@example.com"},
		{"name": "User2", "email": "u2@example.com"},
	}
	for _, u := range users {
		body, _ := json.Marshal(u)
		req := httptest.NewRequest("POST", "/users", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)
		assert.Equal(t, http.StatusCreated, res.Code)
	}

	// Fetch all users
	req := httptest.NewRequest("GET", "/users", nil)
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)

	var allUsers []map[string]interface{}
	err := json.Unmarshal(res.Body.Bytes(), &allUsers)
	assert.NoError(t, err)
	assert.Len(t, allUsers, 2)
}

func TestUpdateUser(t *testing.T) {
	router := setupTestRouter()

	// Create user
	body := map[string]string{"name": "Old Name", "email": "old@example.com"}
	jsonBody, _ := json.Marshal(body)

	createReq := httptest.NewRequest("POST", "/users", bytes.NewBuffer(jsonBody))
	createReq.Header.Set("Content-Type", "application/json")
	createRes := httptest.NewRecorder()
	router.ServeHTTP(createRes, createReq)

	var created map[string]interface{}
	json.Unmarshal(createRes.Body.Bytes(), &created)
	id := int(created["id"].(float64))

	// Update user
	update := map[string]string{"name": "New Name", "email": "new@example.com"}
	updateBody, _ := json.Marshal(update)

	updateReq := httptest.NewRequest("PUT", "/users/"+strconv.Itoa(id), bytes.NewBuffer(updateBody))
	updateReq.Header.Set("Content-Type", "application/json")
	updateRes := httptest.NewRecorder()
	router.ServeHTTP(updateRes, updateReq)

	assert.Equal(t, http.StatusOK, updateRes.Code)

	var updated map[string]interface{}
	json.Unmarshal(updateRes.Body.Bytes(), &updated)
	assert.Equal(t, "New Name", updated["name"])
	assert.Equal(t, "new@example.com", updated["email"])
}

func TestDeleteUser(t *testing.T) {
	router := setupTestRouter()

	// Create user
	body := map[string]string{"name": "Delete Me", "email": "delete@example.com"}
	jsonBody, _ := json.Marshal(body)

	createReq := httptest.NewRequest("POST", "/users", bytes.NewBuffer(jsonBody))
	createReq.Header.Set("Content-Type", "application/json")
	createRes := httptest.NewRecorder()
	router.ServeHTTP(createRes, createReq)

	var created map[string]interface{}
	json.Unmarshal(createRes.Body.Bytes(), &created)
	id := int(created["id"].(float64))

	// Delete user
	deleteReq := httptest.NewRequest("DELETE", "/users/"+strconv.Itoa(id), nil)
	deleteRes := httptest.NewRecorder()
	router.ServeHTTP(deleteRes, deleteReq)

	assert.Equal(t, http.StatusNoContent, deleteRes.Code)

	// Try to get deleted user
	getReq := httptest.NewRequest("GET", "/users/"+strconv.Itoa(id), nil)
	getRes := httptest.NewRecorder()
	router.ServeHTTP(getRes, getReq)

	assert.Equal(t, http.StatusNotFound, getRes.Code)
}
