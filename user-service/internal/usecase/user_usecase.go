package usecase

import (
	"context"
	"errors"
	"log"

	"user-service/internal/models"
	"user-service/internal/repository"
	"user-service/internal/service"
)

var (
	ErrUsernameExists     = errors.New("username already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserNotFound       = errors.New("user not found")
)

type UserUseCase struct {
	repo         repository.UserRepository   // Use interface
	emailService *service.EmailService       // Email service field
}

// Updated constructor accepting interface and email service
func NewUserUseCase(repo repository.UserRepository, emailService *service.EmailService) *UserUseCase {
	return &UserUseCase{
		repo:         repo,
		emailService: emailService,
	}
}

func (u *UserUseCase) RegisterUser(ctx context.Context, user *models.User) error {
	existingUser, err := u.repo.GetUserByUsername(ctx, user.Username)
	if err == nil && existingUser != nil {
		return ErrUsernameExists
	}
	if err != nil && !errors.Is(err, repository.ErrUserNotFound) {
		return err
	}

	if err := u.repo.CreateUser(ctx, user); err != nil {
		return err
	}

	// Send welcome email asynchronously
	go func() {
		if err := u.emailService.SendWelcomeEmail(user.Email, user.Username); err != nil {
			log.Printf("Failed to send welcome email: %v", err)
		}
	}()

	return nil
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
