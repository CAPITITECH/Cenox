package services

import (
	"cenox/database"
	"cenox/models"
	"errors"
)

func CreateUser(user models.User) models.User {
	user.ID = database.GetNextID()
	database.Users = append(database.Users, user)
	return user
}

func GetUsers() []models.User {
	return database.Users
}

func GetUserByID(id int64) (models.User, error) {
	for _, user := range database.Users {
		if user.ID == id {
			return user, nil
		}
	}
	return models.User{}, errors.New("user not found")
}

func UpdateUser(id int64, updatedUser models.User) (models.User, error) {
	for i, user := range database.Users {
		if user.ID == id {
			database.Users[i].Name = updatedUser.Name
			return database.Users[i], nil
		}
	}
	return models.User{}, errors.New("user not found")
}

func DeleteUser(id int64) error {
	for i, user := range database.Users {
		if user.ID == id {
			database.Users = append(database.Users[:i], database.Users[i+1:]...)
			return nil
		}
	}
	return errors.New("user not found")
}