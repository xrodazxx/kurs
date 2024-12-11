package service

import (
	"Kurs/pkg/model"
	"Kurs/pkg/repository"
	"context"
)

type IDevice interface {
	Get(ctx context.Context) ([]model.DeviceData, error)
	Save(ctx context.Context, device model.DeviceData) error
}

type Device struct {
	deviceRepo repository.IDevice
}

func NewDevice(deviceRepo repository.IDevice) *Device {
	return &Device{
		deviceRepo: deviceRepo,
	}
}

func (s *Device) Get(ctx context.Context) ([]model.DeviceData, error) {
	devices, err := s.deviceRepo.GetDevices()
	if err != nil {
		return nil, err
	}

	return devices, nil
}

func (s *Device) Save(ctx context.Context, device model.DeviceData) error {
	if err := s.deviceRepo.SaveDevice(device); err != nil {
		return err
	}

	return nil
}
