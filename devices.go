package kurs

import "time"

// Iot устройства
type DeviceIot struct {
	Id     int    `json:"id"`     // идентификатор устройства
	Name   string `json:"name"`   // название устройства
	Type   string `json:"type"`   // Тип устройства (лампа, колонка,градусник)
	Status string `json:"status"` // Состояние turn(ON or Off)
}

// данные с устройства
type DeviceData struct {
	Id        int       `json:"id"`        // идентификатор записи данных
	DeviceId  int       `json:"deviceId"`  // ссылка на устройство (ID)
	TimeStamp time.Time `json:"timestamp"` // Временная метка
	Value     float64   `json:"value"`     // какие-либо значение от Устройства (Температура-влажность-износ)
}
