package emulator

import (
	"fmt"
	"math/rand"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type DeviceIot struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Status string `json:"status"`
}

type DeviceData struct {
	Id        int       `json:"id"`
	DeviceId  int       `json:"deviceId"`
	TimeStamp time.Time `json:"timestamp"`
	DataType  string    `json:"datatype"`
	Value     float64   `json:"value"`
}

func GenerateData(deviceId int) *DeviceData {
	// Генерация случайных данных для температуры (например)
	return &DeviceData{
		Id:        rand.Int(),
		DeviceId:  deviceId,
		TimeStamp: time.Now(),
		DataType:  "temperature",
		Value:     rand.Float64()*30 + 10, // случайная температура от 10 до 40 градусов
	}
}

func PublishData(client MQTT.Client, deviceId int) {
	data := GenerateData(deviceId)
	topic := fmt.Sprintf("iot/device/%d/data", deviceId)

	// Публикуем данные в топик
	token := client.Publish(topic, 0, false, fmt.Sprintf("%v", data))
	token.Wait()
	fmt.Println("Published data:", data)
}

func StartEmulator(client MQTT.Client, deviceId int) {
	// Эмуляция данных с устройства каждую секунду
	for {
		PublishData(client, deviceId)
		time.Sleep(25 * time.Second)
	}
}
