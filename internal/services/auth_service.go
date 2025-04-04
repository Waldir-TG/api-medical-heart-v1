// internal/services/auth_service.go
package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Waldir-TG/api-medical-heart-v1/internal/models"
	"github.com/Waldir-TG/api-medical-heart-v1/internal/repositories"
	"github.com/Waldir-TG/api-medical-heart-v1/internal/utils"
	"github.com/google/uuid"
)

const (
	// Duración de la sesión: 24 horas
	SessionDuration = 24 * time.Hour
	// Máximo de intentos fallidos antes de bloquear la cuenta
	MaxFailedAttempts = 5
)

type AuthService struct {
	userRepo    *repositories.UserRepository
	sessionRepo *repositories.SessionRepository
}

func NewAuthService(userRepo *repositories.UserRepository, sessionRepo *repositories.SessionRepository) *AuthService {
	return &AuthService{
		userRepo:    userRepo,
		sessionRepo: sessionRepo,
	}
}

func (s *AuthService) Register(ctx context.Context, req *models.RegisterRequest) (uuid.UUID, error) {
	// Verificar si el usuario ya existe
	existingUser, err := s.userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return uuid.Nil, fmt.Errorf("error checking existing user: %w", err)
	}
	if existingUser != nil {
		return uuid.Nil, errors.New("email already registered")
	}

	// Generar hash de la contraseña
	passwordHash, err := utils.HashPassword(req.Password)
	if err != nil {
		return uuid.Nil, fmt.Errorf("error hashing password: %w", err)
	}

	// Crear el usuario
	userID, err := s.userRepo.CreateUser(ctx, req, passwordHash)
	if err != nil {
		return uuid.Nil, fmt.Errorf("error creating user: %w", err)
	}

	return userID, nil
}

func (s *AuthService) Login(ctx context.Context, req *models.LoginRequest, deviceInfo, ipAddress *string) (*models.AuthResponse, error) {
	// Verificar si el usuario existe
	user, err := s.userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, fmt.Errorf("error fetching user: %w", err)
	}
	if user == nil {
		return nil, errors.New("invalid email or password")
	}

	// Verificar si la cuenta está bloqueada por demasiados intentos fallidos
	if user.FailedLoginAttempts >= MaxFailedAttempts {
		return nil, errors.New("account is locked due to too many failed login attempts")
	}

	// Verificar la contraseña
	if !utils.CheckPasswordHash(req.Password, user.PasswordHash) {
		// Incrementar contador de intentos fallidos
		if err := s.userRepo.IncrementFailedLogin(ctx, req.Email); err != nil {
			return nil, fmt.Errorf("error updating failed login attempts: %w", err)
		}
		return nil, errors.New("invalid email or password")
	}

	// Autenticar al usuario utilizando el procedimiento almacenado
	authenticatedUser, err := s.userRepo.AuthenticateUser(ctx, req.Email, user.PasswordHash)
	if err != nil {
		return nil, fmt.Errorf("authentication error: %w", err)
	}
	if authenticatedUser == nil {
		return nil, errors.New("authentication failed")
	}

	// Generar token JWT
	expiresAt := time.Now().Add(SessionDuration)
	token, err := utils.GenerateJWT(authenticatedUser.ID, authenticatedUser.RoleID, authenticatedUser.RoleName, expiresAt)
	if err != nil {
		return nil, fmt.Errorf("error generating JWT: %w", err)
	}

	// Crear sesión en la base de datos
	_, err = s.sessionRepo.CreateSession(ctx, authenticatedUser.ID, token, deviceInfo, ipAddress, expiresAt)
	if err != nil {
		return nil, fmt.Errorf("error creating session: %w", err)
	}

	// Preparar respuesta
	return &models.AuthResponse{
		User: models.User{
			ID:        authenticatedUser.ID,
			Email:     req.Email,
			RoleID:    authenticatedUser.RoleID,
			RoleName:  authenticatedUser.RoleName,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			IsActive:  authenticatedUser.IsActive,
		},
		Token: token,
	}, nil
}

func (s *AuthService) Logout(ctx context.Context, token string) error {
	return s.sessionRepo.InvalidateSession(ctx, token)
}

func (s *AuthService) ValidateToken(ctx context.Context, token string) (*models.Session, *models.User, error) {
	// Verificar el token en la base de datos
	return s.sessionRepo.ValidateSession(ctx, token)
}
