package services

import (
	"context"

	"github.com/Waldir-TG/api-medical-heart-v1/internal/models"
	"github.com/Waldir-TG/api-medical-heart-v1/internal/repositories"
	"github.com/google/uuid"
)

type DeviceService struct {
	deviceRepo *repositories.DeviceRepository
}

func NewDeviceService(deviceRepo *repositories.DeviceRepository) *DeviceService {
	return &DeviceService{
		deviceRepo: deviceRepo,
	}
}

func (s *DeviceService) GetDevices(ctx context.Context) ([]*models.DeviceResponse, error) {
	return s.deviceRepo.GetDevices(ctx)
}

func (s *DeviceService) GetDeviceByID(ctx context.Context, id uuid.UUID) ([]*models.DeviceResponse, error) {
	return s.deviceRepo.GetDeviceByID(ctx, id)
}

func (s *DeviceService) GetDevicesByPatientID(ctx context.Context, patientID uuid.UUID) ([]*models.DeviceResponse, error) {
	return s.deviceRepo.GetDevicesByPatientID(ctx, patientID)
}

func (s *DeviceService) RegisterDevice(ctx context.Context, device *models.DeviceRegisterRequest) (uuid.UUID, error) {
	return s.deviceRepo.RegisterDevice(ctx, device)
}

func (s *DeviceService) UpdateDevice(ctx context.Context, deviceID uuid.UUID, device *models.DeviceUpdateRequest) (string, error) {
	return s.deviceRepo.UpdateDevice(ctx, deviceID, device)
}

func (s *DeviceService) UpdateDeviceSync(ctx context.Context, deviceID uuid.UUID, batteryLevel int) error {
	return s.deviceRepo.UpdateDeviceSync(ctx, deviceID, batteryLevel)
}

func (s *DeviceService) DeactivateDevice(ctx context.Context, deviceID uuid.UUID) (string, error) {
	return s.deviceRepo.DeactivateDevice(ctx, deviceID)
}
