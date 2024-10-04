package services

import (
	"go-fiber-api/internal/models"
	"go-fiber-api/internal/repositories"
)

func GetUsers() ([]models.User, error) {
	return repositories.GetAllUsers()
}

func CreateUser(user *models.User) error {
	return repositories.CreateUser(user)
}

func GetUserByID(id int) (*models.User, error) {
	return repositories.GetUserByID(id)
}

func GetUserByName(name string) ([]models.User, error) {
	return repositories.GetUserByName(name)
}

func UpdateUser(id int, user models.User) (*models.User, error) {
	return repositories.UpdateUser(id, user)
}

func DeleteUser(id int) (*models.User, error) {
	return repositories.DeleteUser(id)
}
