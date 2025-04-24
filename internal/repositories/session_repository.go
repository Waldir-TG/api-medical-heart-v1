// internal/repositories/session_repository.go
package repositories

import (
	"context"
	"errors"
	"time"

	"github.com/Waldir-TG/api-medical-heart-v1/internal/db"
	"github.com/Waldir-TG/api-medical-heart-v1/internal/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type SessionRepository struct {
	db *db.PostgresDB
}

func NewSessionRepository(database *db.PostgresDB) *SessionRepository {
	return &SessionRepository{db: database}
}

func (r *SessionRepository) CreateSession(ctx context.Context, userID uuid.UUID, token string, deviceInfo, ipAddress *string, expiresAt time.Time) (uuid.UUID, error) {
	var sessionID uuid.UUID

	err := r.db.Pool.QueryRow(ctx, `
		CALL create_session($1, $2, $3, $4, $5, NULL)
	`, userID, token, deviceInfo, ipAddress, expiresAt).Scan(&sessionID)

	if err != nil {
		return uuid.Nil, err
	}

	return sessionID, nil
}

func (r *SessionRepository) ValidateSession(ctx context.Context, token string) (*models.Session, *models.User, error) {
	var (
		sessionID uuid.UUID
		userID    uuid.UUID
		roleID    uuid.UUID
		roleName  string
		firstName string
		lastName  string
		isValid   bool
		isExpired bool
	)

	err := r.db.Pool.QueryRow(ctx, `
		SELECT * FROM validate_session($1)
	`, token).Scan(
		&sessionID, &userID, &roleID, &roleName, &firstName, &lastName, &isValid, &isExpired,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil, errors.New("session not found")
		}
		return nil, nil, err
	}

	if !isValid {
		return nil, nil, errors.New("session is not valid")
	}

	if isExpired {
		return nil, nil, errors.New("session has expired")
	}

	session := &models.Session{
		ID:      sessionID,
		UserID:  userID,
		Token:   token,
		IsValid: isValid,
	}

	user := &models.User{
		ID:        userID,
		RoleID:    roleID,
		RoleName:  roleName,
		FirstName: firstName,
		LastName:  lastName,
	}

	return session, user, nil
}

func (r *SessionRepository) InvalidateSession(ctx context.Context, token string) error {
	_, err := r.db.Pool.Exec(ctx, `
		CALL invalidate_session($1)
	`, token)

	return err
}
