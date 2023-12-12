package services

import (
	"backend/internal/core/domain/database"
	"backend/internal/core/domain/payload"
	"backend/internal/core/ports"

	"github.com/google/uuid"
)

type UserService struct {
	userRepository ports.IUserRepository
}

// This line is for get feedback in case we are not implementing the interface correctly
var _ ports.IUserService = (*UserService)(nil)

func NewUserService(repository ports.IUserRepository) *UserService {
	return &UserService{
		userRepository: repository,
	}
}

func (s *UserService) Login(email string, password string) error {
	err := s.userRepository.Login(email, password)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) Register(payload *payload.NewUser) (*database.User, error) {
	user, err := s.userRepository.Register(&database.User{UserId: uuid.New(), Email: payload.Email, Password: payload.Password})
	if err != nil {
		return nil, err
	}
	return user, nil
}
