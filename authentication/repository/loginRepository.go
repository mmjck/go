package repository

import "authentication/models"

type Login struct {
}

var user []models.User

func Init() *Login {
	user = make([]models.User, 0, 2)

	user = append(user,
		models.User{
			Id:       1,
			Name:     "John",
			UserName: "jhon@yopmail.com",
			Password: "jhon123",
			Roles:    []int{1, 2, 3},
		},
		models.User{
			Id:       2,
			Name:     "Maria",
			UserName: "maria@yopmail.com",
			Password: "maria@123",
			Roles:    []int{4},
		},
	)

	return &Login{}
}

func (l *Login) GetUserByName(username, password string) (models.User, *models.ErrorDetail) {
	for _, value := range user {
		if value.UserName == username && value.Password == password {
			return value, nil
		}
	}

	return models.User{}, &models.ErrorDetail{
		ErrorType:    models.ErrorTypeError,
		ErrorMessage: "UserName and password not valid",
	}
}
