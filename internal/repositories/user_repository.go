// internal/repositories/user_repository.go
package repositories

import (
	"context"
	"errors"

	"github.com/Waldir-TG/api-medical-heart-v1/internal/db"
	"github.com/Waldir-TG/api-medical-heart-v1/internal/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type UserRepository struct {
	db *db.PostgresDB
}

func NewUserRepository(database *db.PostgresDB) *UserRepository {
	return &UserRepository{db: database}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *models.RegisterRequest, passwordHash string) (uuid.UUID, error) {
	var userID uuid.UUID
	err := r.db.Pool.QueryRow(ctx, `
		CALL create_user($1, $2, $3, $4, $5, $6, NULL)
	`, user.Email, passwordHash, user.RoleID, user.FirstName, user.LastName, user.PhoneNumber).Scan(&userID)

	if err != nil {
		return uuid.Nil, err
	}

	return userID, nil
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User

	err := r.db.Pool.QueryRow(ctx, `
		SELECT id, email, password_hash, role_id, first_name, last_name, phone_number, 
		       created_at, updated_at, last_login, is_active, failed_login_attempts
		FROM users
		WHERE email = $1
	`, email).Scan(
		&user.ID, &user.Email, &user.PasswordHash, &user.RoleID, &user.FirstName, &user.LastName, &user.PhoneNumber,
		&user.CreatedAt, &user.UpdatedAt, &user.LastLogin, &user.IsActive, &user.FailedLoginAttempts,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) AuthenticateUser(ctx context.Context, email, passwordHash string) (*models.User, error) {
	var user models.User

	row := r.db.Pool.QueryRow(ctx, `
		SELECT * FROM authenticate_user($1, $2)
	`, email, passwordHash)

	err := row.Scan(&user.ID, &user.RoleID, &user.RoleName, &user.IsActive, &user.FailedLoginAttempts)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	// Si el usuario está autenticado pero no está activo, devolvemos un error
	if !user.IsActive {
		return &user, errors.New("user account is not active")
	}

	return &user, nil
}

func (r *UserRepository) IncrementFailedLogin(ctx context.Context, email string) error {
	_, err := r.db.Pool.Exec(ctx, `
		CALL increment_failed_login($1)
	`, email)

	return err
}
