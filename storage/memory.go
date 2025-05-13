package storage

import (
	"devops_project/models"
	"sync"
)

var (
	users   = make(map[int]models.User)
	nextID  = 1
	userMux sync.Mutex
)

func CreateUser(name, email string) models.User {
	userMux.Lock()
	defer userMux.Unlock()

	user := models.User{
		ID:    nextID,
		Name:  name,
		Email: email,
	}
	users[nextID] = user
	nextID++
	return user
}

func GetAllUsers() []models.User {
	userMux.Lock()
	defer userMux.Unlock()

	result := make([]models.User, 0, len(users))
	for _, u := range users {
		result = append(result, u)
	}
	return result
}

func GetUser(id int) (models.User, bool) {
	userMux.Lock()
	defer userMux.Unlock()

	user, found := users[id]
	return user, found
}

func UpdateUser(id int, name, email string) (models.User, bool) {
	userMux.Lock()
	defer userMux.Unlock()

	if _, exists := users[id]; !exists {
		return models.User{}, false
	}
	user := models.User{
		ID:    id,
		Name:  name,
		Email: email,
	}
	users[id] = user
	return user, true
}

func DeleteUser(id int) bool {
	userMux.Lock()
	defer userMux.Unlock()

	if _, exists := users[id]; !exists {
		return false
	}
	delete(users, id)
	return true
}

func Reset() {
	userMux.Lock()
	defer userMux.Unlock()
	users = make(map[int]models.User)
	nextID = 1
}
