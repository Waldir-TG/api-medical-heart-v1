package repositories

import (
	"context"
	"fmt"

	"github.com/Waldir-TG/api-medical-heart-v1/internal/db"
	"github.com/Waldir-TG/api-medical-heart-v1/internal/models"
	"github.com/google/uuid"
)

type PatientRepo struct {
	db *db.PostgresDB
}

func NewPatientRepo(db *db.PostgresDB) *PatientRepo {
	return &PatientRepo{db}
}

func (r *PatientRepo) GetPatientsBasicDetail(ctx context.Context, patientID *uuid.UUID) ([]*models.PatientDetailBasicResponse, error) {
	query := `SELECT * FROM get_patients_basic_details($1);`
	fmt.Println(patientID)

	rows, err := r.db.Pool.Query(ctx, query, patientID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var patients []*models.PatientDetailBasicResponse

	for rows.Next() {
		var patient models.PatientDetailBasicResponse
		if err := rows.Scan(
			&patient.PatientID,
			&patient.UserID,
			&patient.Email,
			&patient.FirstName,
			&patient.LastName,
			&patient.PhoneNumber,
			&patient.IsActive,
			&patient.DateOfBirth,
			&patient.Gender,
			&patient.BloodType,
			&patient.HeightCm,
			&patient.WeightKg,
			&patient.MedicalConditions,
			&patient.Allergies,
			&patient.EmergencyContactName,
			&patient.EmergencyContactPhone,
			&patient.EmergencyContactRelation,
			&patient.MinHeartRate,
			&patient.MaxHeartRate,
			&patient.MonitoringActive,
			&patient.CreatedAt,
			&patient.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}
		patients = append(patients, &patient)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration failed: %w", err)
	}

	return patients, nil
}

func (r *PatientRepo) GetPatientsDetails(ctx context.Context, patientID *uuid.UUID) ([]*models.PatientDetailResponse, error) {
	query := `SELECT * FROM get_patients_details($1);`
	fmt.Println(patientID)
	rows, err := r.db.Pool.Query(ctx, query, patientID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var patients []*models.PatientDetailResponse

	for rows.Next() {
		var patient models.PatientDetailResponse
		if err := rows.Scan(
			&patient.PatientID,
			&patient.UserID,
			&patient.Email,
			&patient.FirstName,
			&patient.LastName,
			&patient.PhoneNumber,
			&patient.IsActive,
			&patient.DateOfBirth,
			&patient.Gender,
			&patient.BloodType,
			&patient.HeightCm,
			&patient.WeightKg,
			&patient.MedicalConditions,
			&patient.Allergies,
			&patient.EmergencyContactName,
			&patient.EmergencyContactPhone,
			&patient.EmergencyContactRelation,
			&patient.MinHeartRate,
			&patient.MaxHeartRate,
			&patient.MonitoringActive,
			&patient.CreatedAt,
			&patient.UpdatedAt,
			&patient.AssignedDoctors,
			&patient.FamilyMembers,
			&patient.LastHeartReadings,
		); err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}
		patients = append(patients, &patient)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration failed: %w", err)
	}
	return patients, nil
}

func (r *PatientRepo) CreatePatient(
	ctx context.Context,
	patient *models.PatientCreateRequest) (uuid.UUID, error) {
	query := `CALL create_patient(NULL,$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13);`
	var patientID uuid.UUID
	err := r.db.Pool.QueryRow(ctx, query,
		patient.UserID,
		patient.DateOfBirth,
		patient.Gender,
		patient.BloodType,
		patient.HeightCm,
		patient.WeightKg,
		patient.MedicalConditions,
		patient.Allergies,
		patient.EmergencyContactName,
		patient.EmergencyContactPhone,
		patient.EmergencyContactRelation,
		patient.MinHeartRate,
		patient.MaxHeartRate,
	).Scan(&patientID)
	if err != nil {
		return uuid.Nil, err
	}

	return patientID, nil
}

func (r *PatientRepo) UpdatePatient(ctx context.Context, patientID uuid.UUID, patient *models.PatientUpdatedRequest) (string, error) {
	query := `CALL update_patient_info($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12);`
	_, err := r.db.Pool.Exec(ctx, query,
		patientID,
		// patient.DateOfBirth,
		// patient.Gender,
		// patient.BloodType,
		patient.HeightCm,
		patient.WeightKg,
		patient.MedicalConditions,
		patient.Allergies,
		patient.EmergencyContactName,
		patient.EmergencyContactPhone,
		patient.EmergencyContactRelation,
		patient.MinHeartRate,
		patient.MaxHeartRate,
		patient.MonitoringActive,
		patient.AlertRecipients,
	)
	if err != nil {
		return "", err
	}

	return "Patient updated successfully", nil
}

func (r *PatientRepo) AssignDoctorToPatient(ctx context.Context, DoctorPatientAssign *models.DoctorPatientAssignRequest) (string, error) {
	query := `CALL assign_doctor_to_patient($1, $2);`
	_, err := r.db.Pool.Exec(ctx, query,
		DoctorPatientAssign.DoctorID,
		DoctorPatientAssign.PatientID,
	)
	if err != nil {
		return "", err
	}

	return "Doctor assigned to patient successfully", nil
}

// func (r *PatientRepo) AssignFamilyMemberToPatient(ctx context.Context, familyMember *models.FamilyMemberAssignRequest) (string, error) {

// }
