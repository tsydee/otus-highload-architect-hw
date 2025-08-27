package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/tsydim/otus-highload-architect-hw/internal/apperrs"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/tsydim/otus-highload-architect-hw/internal/config"
	"github.com/tsydim/otus-highload-architect-hw/internal/users"
)

const tokenTTL = 1 * time.Hour

type AuthService struct {
	users     *users.UserService
	passwords *users.PasswordService
	cfg       config.Security
}

func NewAuthService(
	users *users.UserService,
	passwords *users.PasswordService,
	cfg config.Security,
) *AuthService {
	return &AuthService{
		users:     users,
		passwords: passwords,
		cfg:       cfg,
	}
}

type SignUpData = users.UserServiceCreate

func (s *AuthService) SignUp(ctx context.Context, payload SignUpData) (string, error) {
	user, err := s.users.Create(ctx, payload)
	if err != nil {
		return "", fmt.Errorf("create new user: %w", err)
	}

	token, err := s.generateToken(user.ID)
	if err != nil {
		return "", fmt.Errorf("generate token: %w", err)
	}

	return token, nil
}

func (s *AuthService) SignIn(ctx context.Context, id string, password string) (string, error) {
	user, err := s.users.Get(ctx, id)
	if err != nil {
		return "", fmt.Errorf("signin: %w", err)
	}
	if !user.Password.IsSamePassword(password) {
		return "", apperrs.ErrUnauthorize
	}
	token, err := s.generateToken(user.ID)
	if err != nil {
		return "", fmt.Errorf("generate token: %w", err)
	}
	return token, nil
}

func (s *AuthService) Verify(token string) (users.UserID, error) {
	var zero users.UserID
	token = strings.TrimSpace(token)

	t, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(s.cfg.SecretKey), nil
	})
	if err != nil {
		return zero, errors.Join(apperrs.ErrUnauthorize, fmt.Errorf("parse token: %w", err))
	}

	claims := t.Claims.(*jwt.StandardClaims)
	return claims.Issuer, nil
}

func (s *AuthService) generateToken(userID users.UserID) (string, error) {
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(tokenTTL).Unix(),
		Issuer:    userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.cfg.SecretKey))
}
