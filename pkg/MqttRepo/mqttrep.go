package mqttrepo

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type DeviceData struct {
	DeviceID  string  `db:"device_id"`
	Timestamp string  `db:"timestamp"`
	DataType  string  `db:"data_type"`
	Value     float64 `db:value`
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
func (d *DB) saveDevice(data DeviceData) error {
	query := "INSERT INTO device_data (device_id, timestamp, data_type, value) VALUES ($1, $2, $3, $4)"
	_, err := d.db.Exec(query, data.DeviceID, data.Timestamp, data.DataType, data.Value)
	if err != nil {
		return fmt.Errorf("Невозможно сохранить устройство: %v", err)
	}
	return nil
}
