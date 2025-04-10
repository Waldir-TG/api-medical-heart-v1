package repositories

import (
	"context"
	"fmt"

	"github.com/Waldir-TG/api-medical-heart-v1/internal/db"
	"github.com/Waldir-TG/api-medical-heart-v1/internal/models"
)

type DoctorRepository struct {
	db *db.PostgresDB
}

func NewDoctorRepository(database *db.PostgresDB) *DoctorRepository {
	return &DoctorRepository{db: database}
}

func (r *DoctorRepository) GetDoctors(ctx context.Context) ([]*models.Doctor, error) {

	query := `SELECT * FROM doctors`
	rows, err := r.db.Pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	var doctors []*models.Doctor
	for rows.Next() {
		var doctor models.Doctor
		if err := rows.Scan(
			&doctor.ID,
			&doctor.IdentityType,
			&doctor.IdentityNumber,
			&doctor.Specialty,
			&doctor.LicenseNumber,
			&doctor.HospitalAffiliaton,
			&doctor.CreatedAt,
			&doctor.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}
		doctors = append(doctors, &doctor)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}
	return doctors, nil
}
