package interfaces

import (
	"register/models/domain"
	"register/models/dtos"
)

type AuthService interface {
	Register(request dtos.RegisterRequest) (domain.User, error)
}
