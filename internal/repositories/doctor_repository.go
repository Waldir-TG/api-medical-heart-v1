package repositories

import (
	"context"
	"fmt"

	"github.com/Waldir-TG/api-medical-heart-v1/internal/db"
	"github.com/Waldir-TG/api-medical-heart-v1/internal/models"
	"github.com/google/uuid"
)

type DoctorRepository struct {
	db *db.PostgresDB
}

func NewDoctorRepository(database *db.PostgresDB) *DoctorRepository {
	return &DoctorRepository{db: database}
}

func (r *DoctorRepository) GetDoctors(ctx context.Context) ([]*models.DoctorDetailsResponse, error) {

	query := `SELECT * FROM get_all_doctors()`
	rows, err := r.db.Pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var doctors []*models.DoctorDetailsResponse
	for rows.Next() {
		var doctor models.DoctorDetailsResponse
		if err := rows.Scan(
			&doctor.DoctorID,
			&doctor.UserID,
			&doctor.Email,
			&doctor.FirstName,
			&doctor.LastName,
			&doctor.PhoneNumber,
			&doctor.IdentityType,
			&doctor.IdentityNumber,
			&doctor.Specialty,
			&doctor.LicenseNumber,
			&doctor.HospitalAffiliaton,
			&doctor.CreatedAt,
			&doctor.UpdatedAt,
			&doctor.IsActive,
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

func (r *DoctorRepository) GetDoctorDetail(ctx context.Context, id uuid.UUID) (*models.DoctorDetailsResponse, error) {

	query := `SELECT * FROM get_doctor_details($1);`
	row := r.db.Pool.QueryRow(ctx, query, id)
	var doctor models.DoctorDetailsResponse
	if err := row.Scan(
		&doctor.DoctorID,
		&doctor.UserID,
		&doctor.Email,
		&doctor.FirstName,
		&doctor.LastName,
		&doctor.PhoneNumber,
		&doctor.IdentityType,
		&doctor.IdentityNumber,
		&doctor.Specialty,
		&doctor.LicenseNumber,
		&doctor.HospitalAffiliaton,
		&doctor.CreatedAt,
		&doctor.UpdatedAt,
		&doctor.IsActive,
	); err != nil {
		return nil, fmt.Errorf("scan failed: %w", err)
	}
	return &doctor, nil
}

func (r *DoctorRepository) CreateDoctor(ctx context.Context, doctor *models.DoctorCreateRequest) (uuid.UUID, error) {
	var doctorID uuid.UUID
	query := `CALL create_doctor(NULL,$1, $2, $3, $4, $5, $6);`
	err := r.db.Pool.QueryRow(ctx, query, doctor.UserID, doctor.IdentityType, doctor.IdentityNumber, doctor.Specialty, doctor.LicenseNumber, doctor.HospitalAffiliation).Scan(&doctorID)
	if err != nil {
		return uuid.Nil, err
	}
	return doctorID, nil
}

func (r *DoctorRepository) UpdateDoctor(ctx context.Context, doctorID uuid.UUID, doctor *models.DoctorUpdateRequest) (string, error) {

	query := `CALL update_doctor_info($1, $2, $3, $4, $5, $6, $7);`
	_, err := r.db.Pool.Exec(ctx, query, doctorID, doctor.IdentityType, doctor.IdentityNumber, doctor.Specialty, doctor.LicenseNumber, doctor.HospitalAffiliation, doctor.IsActive)
	if err != nil {
		return "", fmt.Errorf("error updated doctor: %w", err)
	}
	return "Actualizado exitosamente", nil
}
