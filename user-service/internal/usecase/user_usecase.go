package usecase

import (
	"context"
	"errors"

	"user-service/internal/models"
	"user-service/internal/repository"
)

var (
	ErrUsernameExists      = errors.New("username already exists")
	ErrInvalidCredentials  = errors.New("invalid credentials")
	ErrUserNotFound        = errors.New("user not found")
)

type UserUseCase struct {
	repo repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) *UserUseCase {
	return &UserUseCase{repo: repo}
}

func (u *UserUseCase) RegisterUser(ctx context.Context, user *models.User) error {
	existingUser, err := u.repo.GetUserByUsername(ctx, user.Username)
	if err == nil && existingUser != nil {
		return ErrUsernameExists
	}
	if err != nil && !errors.Is(err, repository.ErrUserNotFound) {
		return err
	}

	return u.repo.CreateUser(ctx, user)
}

func (u *UserUseCase) AuthenticateUser(ctx context.Context, username, password string) (*models.User, error) {
	user, err := u.repo.GetUserByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return nil, ErrInvalidCredentials
		}
		return nil, err
	}

	if user.Password != password {
		return nil, ErrInvalidCredentials
	}

	return user, nil
}

func (u *UserUseCase) GetUserProfile(ctx context.Context, id string) (*models.User, error) {
	user, err := u.repo.GetUserByID(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return user, nil
}