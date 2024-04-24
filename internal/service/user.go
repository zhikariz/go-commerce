package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/zhikariz/go-commerce/internal/entity"
	"github.com/zhikariz/go-commerce/internal/repository"
	"github.com/zhikariz/go-commerce/pkg/token"
)

type UserService interface {
	Login(email string, password string) (string, error)
	FindAllUser() ([]entity.User, error)
}

type userService struct {
	userRepository repository.UserRepository
	tokenUseCase   token.TokenUseCase
}

func NewUserService(userRepository repository.UserRepository, tokenUseCase token.TokenUseCase) *userService {
	return &userService{
		userRepository: userRepository,
		tokenUseCase:   tokenUseCase,
	}
}

func (s *userService) Login(email string, password string) (string, error) {
	user, err := s.userRepository.FindUserByEmail(email)
	if err != nil {
		return "", errors.New("email/password yang anda masukkan salah")
	}

	if user.Password != password {
		return "", errors.New("email/password yang anda masukkan salah")
	}

	expiredTime := time.Now().Local().Add(5 * time.Minute)

	claims := token.JwtCustomClaims{
		ID:     user.ID.String(),
		Email:  user.Email,
		Role:   user.Role,
		Alamat: user.Alamat,
		NoHP:   user.NoHp,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Go-Commerce",
			ExpiresAt: jwt.NewNumericDate(expiredTime),
		},
	}

	token, err := s.tokenUseCase.GenerateAccessToken(claims)
	if err != nil {
		return "", errors.New("ada kesalahan dari sistem")
	}

	return token, nil
}

func (s *userService) FindAllUser() ([]entity.User, error) {
	return s.userRepository.FindAllUser()
}
