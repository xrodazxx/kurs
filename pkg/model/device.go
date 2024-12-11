package model

type DeviceData struct {
	DeviceID   string  `db:"device_id"`
	DeviceName string  `db:"device_name"`
	Timestamp  string  `db:"timestamp"`
	DataType   string  `db:"data_type"`
	Value      float64 `db:"value"`
}
