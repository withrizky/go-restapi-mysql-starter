package repository

import (
	"myapp/internal/database"
	"myapp/internal/model"
)

func CreateUser(user *model.User) error {
	return database.DB.Create(user).Error
}
