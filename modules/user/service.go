package user

import "go-api-rest/modules/database"

func GetUsers() (users []User) {
	database.Connection.Find(&users)
	return users
}

func CreateUser(user User) (*User, error) {
	result := database.Connection.Create(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
