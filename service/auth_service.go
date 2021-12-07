package service

import (
	"register/models/domain"
	"register/models/dtos"
	"register/models/interfaces"

	"gorm.io/gorm"
)

type AuthService struct {
	db *gorm.DB
}

func NewAuthService(db *gorm.DB) interfaces.AuthService {
	return &AuthService{db: db}
}

func (s AuthService) Register(request dtos.RegisterRequest) (domain.User, error) {

	user := domain.User{
		Username: request.Username,
		Password: request.Password,
	}

	// Create a new user with the given username and password and save it to the database
	err := s.db.Create(&user).Error
	// Check for errors and return the error if the user could not be created
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}
