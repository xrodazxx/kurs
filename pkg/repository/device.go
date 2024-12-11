package repository

import (
	"Kurs/pkg/model"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type IDevice interface {
	SaveDevice(data model.DeviceData) error
	GetDevices() ([]model.DeviceData, error)
}

type Device struct {
	db *sqlx.DB
}

func NewDevice(db *sqlx.DB) *Device {
	return &Device{db: db}
}

func (d *Device) SaveDevice(data model.DeviceData) error {
	query := `
		INSERT INTO device_data (device_id, device_name, timestamp, data_type, value) VALUES ($1, $2, NOW(), $3, $4)`
	_, err := d.db.Exec(query, data.DeviceID, data.DeviceName, data.DataType, data.Value)
	if err != nil {
		return fmt.Errorf("Невозможно сохранить устройство: %v", err)
	}
	return nil
}

func (d *Device) GetDevices() ([]model.DeviceData, error) {
	var devices []model.DeviceData
	query := `
		SELECT device_id, device_name, timestamp, data_type, value 
		FROM device_data`

	err := d.db.Select(&devices, query)
	if err != nil {
		return nil, fmt.Errorf("Ошибка при получении списка устройств: %v", err)
	}
	return devices, nil
}
