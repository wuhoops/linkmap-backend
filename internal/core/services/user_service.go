package services

import (
	"backend/internal/core/domain/database"
	"backend/internal/core/ports"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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

func (s *UserService) GetUserById(userId string) (*database.User, error) {
	user, err := s.userRepository.GetUserById(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) Login(email string, password string) error {
	err := s.userRepository.Login(email, password)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) Register(payload *database.User) error {
	payload.UserId = uuid.New().String()
	bytes, err := bcrypt.GenerateFromPassword([]byte(payload.UserId), 14)
	if err != nil {
		return err
	}
	payload.Password = string(bytes)

	if err := s.userRepository.Register(payload); err != nil {
		return err
	}
	return nil
}

func (s *UserService) CreateUserName(userId string, userName string) error {
	err := s.userRepository.CreateUserName(userId, userName)
	if err != nil {
		return err
	}
	return nil
}
