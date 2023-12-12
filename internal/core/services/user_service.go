package services

import (
	"backend/internal/core/domain/database"
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

func (s *UserService) Register(payload *database.User) error {
	payload.UserId = uuid.New()
	err := s.userRepository.Register(&database.User{UserId: payload.UserId, Email: payload.Email, Password: payload.Password})
	if err != nil {
		return err
	}
	return nil
}
