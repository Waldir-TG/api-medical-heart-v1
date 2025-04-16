package repositories

import (
	"context"

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

//	type PatientCreateRequest struct {
//		UserID                   string   `json:"user_id" validate:"required,uuid4"`
//		DateOfBirth              string   `json:"date_of_birth" validate:"required,datetime=2006-01-02"`
//		Gender                   string   `json:"gender" validate:"required,oneof=male female other"`
//		BloodType                string   `json:"blood_type" validate:"omitempty,oneof=A+ A- B+ B- AB+ AB- O+ O-"`
//		HeightCm                 int      `json:"height_cm" validate:"required,min=30,max=300"`
//		WeightKg                 float64  `json:"weight_kg" validate:"required,min=0.5,max=500"`
//		MedicalConditions        []string `json:"medical_conditions"`
//		Allergies                []string `json:"allergies"`
//		EmergencyContactName     string   `json:"emergency_contact_name" validate:"required"`
//		EmergencyContactPhone    string   `json:"emergency_contact_phone" validate:"required"`
//		EmergencyContactRelation string   `json:"emergency_contact_relation" validate:"required"`
//		MinHeartRate             int      `json:"min_heart_rate" validate:"required,min=20,max=200"`
//		MaxHeartRate             int      `json:"max_heart_rate" validate:"required,min=20,max=200,gtfield=MinHeartRate"`
//	}
func (r *PatientRepo) CreatePatient(
	ctx context.Context,
	patient *models.PatientCreateRequest) (uuid.UUID, error) {
	query := `CALL create_doctor(NULL,$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13,$14);`
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
	query := `CALL update_patient($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14);`
	_, err := r.db.Pool.Exec(ctx, query,
		patientID,
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
	)
	if err != nil {
		return "", err
	}

	return "Patient updated successfully", nil
}
