package users

import (
	"context"
	"fmt"
)

type repository interface {
	Create(ctx context.Context, user User) error
	Get(ctx context.Context, userID UserID) (User, error)
}

type UserService struct {
	repo      repository
	passwords *PasswordService
}

func NewUserService(repo repository, passwordService *PasswordService) *UserService {
	return &UserService{
		repo:      repo,
		passwords: passwordService,
	}
}

type UserServiceCreate struct {
	UserFields
	Password string `json:"password"`
}

func (s *UserService) Create(ctx context.Context, userData UserServiceCreate) (User, error) {
	var zero User
	passwordHash, err := s.passwords.FromString(userData.Password)
	if err != nil {
		return zero, fmt.Errorf("prepare password: %w", err)
	}
	user, err := NewUser(
		userData.UserFields,
		passwordHash,
	)
	if err != nil {
		return zero, fmt.Errorf("create new user: %w", err)
	}
	err = s.repo.Create(ctx, user)
	if err != nil {
		return zero, fmt.Errorf("persist new user: %w", err)
	}

	return user, nil
}

func (s *UserService) Get(ctx context.Context, userID UserID) (User, error) {
	return s.repo.Get(ctx, userID)
}
