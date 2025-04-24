package repositories

import (
	"context"
	"fmt"

	"github.com/Waldir-TG/api-medical-heart-v1/internal/db"
	"github.com/Waldir-TG/api-medical-heart-v1/internal/models"
	"github.com/google/uuid"
)

type DeviceRepository struct {
	db *db.PostgresDB
}

func NewDeviceRepository(database *db.PostgresDB) *DeviceRepository {
	return &DeviceRepository{db: database}
}

func (r *DeviceRepository) GetDevices(ctx context.Context) ([]*models.DeviceResponse, error) {
	query := `SELECT * FROM get_devices($1,$2)`
	rows, err := r.db.Pool.Query(ctx, query, nil, nil)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var devices []*models.DeviceResponse
	for rows.Next() {
		var device models.DeviceResponse

		if err := rows.Scan(
			&device.DeviceID,
			&device.PatientID,
			&device.DeviceType,
			&device.SerialNumber,
			&device.FirmwareVersion,
			&device.LastSync,
			&device.BatteryLevel,
			&device.IsActive,
			&device.RegisteredAt,
		); err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}

		devices = append(devices, &device)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}
	return devices, nil
}

func (r *DeviceRepository) GetDeviceByID(ctx context.Context, id uuid.UUID) ([]*models.DeviceResponse, error) {
	query := `SELECT * FROM get_devices($1, $2);`
	fmt.Println(id)
	rows, err := r.db.Pool.Query(ctx, query, id, nil)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var devices []*models.DeviceResponse
	for rows.Next() {
		var device models.DeviceResponse

		if err := rows.Scan(
			&device.DeviceID,
			&device.PatientID,
			&device.DeviceType,
			&device.SerialNumber,
			&device.FirmwareVersion,
			&device.LastSync,
			&device.BatteryLevel,
			&device.IsActive,
			&device.RegisteredAt,
		); err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}

		devices = append(devices, &device)
	}
	fmt.Println(devices)
	return devices, nil
}

func (r *DeviceRepository) GetDevicesByPatientID(ctx context.Context, patientID uuid.UUID) ([]*models.DeviceResponse, error) {
	query := `SELECT * FROM get_devices($1,$2);`
	rows, err := r.db.Pool.Query(ctx, query, nil, patientID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var devices []*models.DeviceResponse
	for rows.Next() {
		var device models.DeviceResponse

		if err := rows.Scan(
			&device.DeviceID,
			&device.PatientID,
			&device.DeviceType,
			&device.SerialNumber,
			&device.FirmwareVersion,
			&device.LastSync,
			&device.BatteryLevel,
			&device.IsActive,
			&device.RegisteredAt,
		); err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}

		devices = append(devices, &device)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}
	return devices, nil
}

func (r *DeviceRepository) RegisterDevice(ctx context.Context, device *models.DeviceRegisterRequest) (uuid.UUID, error) {
	var deviceID uuid.UUID
	var patientUUID *uuid.UUID

	// Convert patient ID string to UUID if provided
	if device.PatientID != "" {
		parsedUUID, err := uuid.Parse(device.PatientID)
		if err != nil {
			return uuid.Nil, fmt.Errorf("invalid patient ID format: %w", err)
		}
		patientUUID = &parsedUUID
	}

	query := `CALL register_device(NULL, $1, $2, $3, $4);`
	err := r.db.Pool.QueryRow(ctx, query,
		patientUUID,
		device.DeviceType,
		device.SerialNumber,
		device.FirmwareVersion,
	).Scan(&deviceID)

	if err != nil {
		return uuid.Nil, err
	}
	return deviceID, nil
}

func (r *DeviceRepository) UpdateDevice(ctx context.Context, deviceID uuid.UUID, device *models.DeviceUpdateRequest) (string, error) {

	query := `CALL update_device($1, $2, $3, $4, $5, $6);`
	_, err := r.db.Pool.Exec(ctx, query,
		deviceID,
		device.PatientID,
		device.DeviceType,
		device.FirmwareVersion,
		device.BatteryLevel,
		device.IsActive,
	)

	if err != nil {
		return "", fmt.Errorf("error updating device: %w", err)
	}
	return "Device updated successfully", nil
}

func (r *DeviceRepository) UpdateDeviceSync(ctx context.Context, deviceID uuid.UUID, batteryLevel int) error {
	query := `CALL update_device_sync($1, $2);`
	_, err := r.db.Pool.Exec(ctx, query, deviceID, batteryLevel)
	if err != nil {
		return fmt.Errorf("error updating device sync: %w", err)
	}
	return nil
}

func (r *DeviceRepository) DeactivateDevice(ctx context.Context, deviceID uuid.UUID) (string, error) {
	query := `CALL deactivate_device($1);`
	_, err := r.db.Pool.Exec(ctx, query, deviceID)
	if err != nil {
		return "", fmt.Errorf("error deactivating device: %w", err)
	}
	return "Device deactivated successfully", nil
}
