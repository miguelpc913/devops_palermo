package repository

import (
	"devops_project/db/models"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) CreateUser(name, email string) error {
	user := models.User{
		Name:  name,
		Email: email,
	}
	return r.db.Create(&user).Error
}

func (r *Repository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := r.db.Find(&users)
	return users, result.Error
}

func (r *Repository) GetUser(id int) (models.User, error) {
	var user models.User
	result := r.db.First(&user, id)
	return user, result.Error
}

func (r *Repository) UpdateUser(id int, name, email string) (models.User, error) {
	var user models.User
	result := r.db.First(&user, id)
	if result.Error != nil {
		return models.User{}, result.Error
	}

	user.Name = name
	user.Email = email

	saveResult := r.db.Save(&user)
	return user, saveResult.Error
}

func (r *Repository) DeleteUser(id int) error {
	result := r.db.Delete(&models.User{}, id)
	return result.Error
}
