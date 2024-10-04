package repositories

import (
	"errors"
	"go-fiber-api/database"
	"go-fiber-api/internal/models"

	"gorm.io/gorm"
)

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := database.DB.Find(&users)
	return users, result.Error
}

func CreateUser(user *models.User) error {
	result := database.DB.Create(user)
	return result.Error
}

func GetUserByID(id int) (*models.User, error) {
	var users models.User
	result := database.DB.First(&users, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &users, nil
}

func GetUserByName(name string) ([]models.User, error) {
	var users []models.User
	result := database.DB.Where("name = ?", name).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, result.Error
}

func UpdateUser(id int, updateUser models.User) (*models.User, error) {
	var user models.User
	result := database.DB.First(&user, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}

	user.Name = updateUser.Name
	user.Email = updateUser.Email

	result = database.DB.Save(&user)
	return &user, result.Error
}

func DeleteUser(id int) (*models.User, error) {
	var user models.User
	result := database.DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}

	result = database.DB.Delete(&user)
	return &user, result.Error
}
