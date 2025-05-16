package db_utils

import (
	"devops_project/db/models"
	"fmt"

	"gorm.io/gorm"
)

func InitDB() (db *gorm.DB, err error) {
	db, err = ConnectDb()
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&models.User{})
	fmt.Println("Successfully connected!", db)
	return db, nil
}
