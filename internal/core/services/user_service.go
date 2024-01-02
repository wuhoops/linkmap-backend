package services

import (
	"backend/internal/core/domain/database"
	"backend/internal/core/ports"
	"backend/internal/util/config"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
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

func (s *UserService) Login(payload *database.User) error {
	password := payload.Password
	err := s.userRepository.Login(payload)
	if err != nil {
		return err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(payload.Password), []byte(password)); err != nil {
		return err
	}
	return nil
}

func (s *UserService) Register(payload *database.User) error {
	payload.UserId = uuid.New().String()
	bytes, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
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

func (s *UserService) GetUserByUsername(userName string) (*database.User, error) {
	user, err := s.userRepository.GetUserByUsername(userName)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) SetRefreshToken(username string, refreshToken string, expiration time.Duration) error {
	key := "refresh_" + username
	err := s.userRepository.SetRefreshToken(key, refreshToken, expiration)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) GetRefreshToken(username string) (string, error) {
	key := "refresh_" + username
	refreshToken, err := s.userRepository.GetRefreshToken(key)
	if err != nil {
		return "", err
	}
	return refreshToken, nil
}

func (s *UserService) GenerateToken(username string, expiration time.Time) (string, error) {
	privateKey := []byte(config.C.Secret)
	claims := jwt.MapClaims{
		"name": username,
		"exp":  expiration.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		log.Printf("token.SignedString: %v", err)
		return "", err
	}
	return tokenString, nil

}
