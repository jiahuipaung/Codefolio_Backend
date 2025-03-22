package app

import (
	"context"
	"errors"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/jiahuipaung/Codefolio_Backend/user/domain"
)

var (
	ErrUserExists         = errors.New("user already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type UserService struct {
	userRepo  domain.UserRepository
	jwtSecret []byte
}

func NewUserService(userRepo domain.UserRepository, jwtSecret string) *UserService {
	return &UserService{
		userRepo:  userRepo,
		jwtSecret: []byte(jwtSecret),
	}
}

type RegisterRequest struct {
	Username string
	Password string
	Email    string
}

type LoginRequest struct {
	Username string
	Password string
}

type TokenResponse struct {
	Token     string
	ExpiresAt time.Time
}

func (s *UserService) Register(ctx context.Context, req RegisterRequest) error {
	// 检查用户名是否已存在
	if _, err := s.userRepo.FindByUsername(ctx, req.Username); err == nil {
		return ErrUserExists
	}

	// 检查邮箱是否已存在
	if _, err := s.userRepo.FindByEmail(ctx, req.Email); err == nil {
		return ErrUserExists
	}

	user, err := domain.NewUser(req.Username, req.Password, req.Email)
	if err != nil {
		return err
	}

	return s.userRepo.Create(ctx, user)
}

func (s *UserService) Login(ctx context.Context, req LoginRequest) (*TokenResponse, error) {
	user, err := s.userRepo.FindByUsername(ctx, req.Username)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	if !user.ValidatePassword(req.Password) {
		return nil, ErrInvalidCredentials
	}

	// 生成 JWT token
	expiresAt := time.Now().Add(24 * time.Hour)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     expiresAt.Unix(),
	})

	tokenString, err := token.SignedString(s.jwtSecret)
	if err != nil {
		return nil, err
	}

	return &TokenResponse{
		Token:     tokenString,
		ExpiresAt: expiresAt,
	}, nil
}
