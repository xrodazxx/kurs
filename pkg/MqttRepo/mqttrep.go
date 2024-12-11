package mqttrepo

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type DeviceData struct {
	DeviceID  string    `db:"device_id"`
	Timestamp time.Time `db:"timestamp"`
	DataType  string    `db:"data_type"`
	Value     float64   `db:"value"`
}

type DB struct {
	db *sqlx.DB
}

func NewDB(iotName string) (*DB, error) {
	db, err := sqlx.Connect("postgres", iotName)
	if err != nil {
		return nil, err
	}
	return &DB{db: db}, nil
}

func (d *DB) SaveDevice(data DeviceData) error {
	query := `
		INSERT INTO device_data (device_id, timestamp, data_type, value)
		VALUES (:device_id, :timestamp, :data_type, :value)`
	_, err := d.db.NamedExec(query, data)
	if err != nil {
		return fmt.Errorf("Невозможно сохранить устройство: %v", err)
	}
	return nil
}
