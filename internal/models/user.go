// internal/models/user.go
package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID                  uuid.UUID  `json:"id"`
	Email               string     `json:"email"`
	PasswordHash        string     `json:"-"` // No se serializa en JSON
	RoleID              uuid.UUID  `json:"role_id"`
	RoleName            string     `json:"role_name,omitempty"`
	FirstName           string     `json:"first_name"`
	LastName            string     `json:"last_name"`
	PhoneNumber         *string    `json:"phone_number,omitempty"`
	CreatedAt           time.Time  `json:"created_at"`
	UpdatedAt           time.Time  `json:"updated_at"`
	LastLogin           *time.Time `json:"last_login,omitempty"`
	IsActive            bool       `json:"is_active"`
	FailedLoginAttempts int        `json:"-"`
	ResetToken          *string    `json:"-"`
	ResetTokenExpires   *time.Time `json:"-"`
}

// PasswordResetRequest representa la solicitud de reseteo de contraseña
type PasswordResetRequest struct {
	Email string `json:"email" validate:"required,email"`
}

// PasswordUpdateRequest representa la solicitud para actualizar contraseña
type PasswordUpdateRequest struct {
	OldPassword string `json:"old_password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required,min=8"`
}

// UserUpdateRequest representa la solicitud para actualizar un usuario
type UserUpdateRequest struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
}
