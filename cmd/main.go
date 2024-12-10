package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	// Подключение к базе данных PostgreSQL
	db, err := repository.NewDB("user=postgres password=sosi dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.db.Close()

	// Инициализация MQTT клиента
	mqttClient := mqtt.InitMQTT("broker.emqx.io", "1883")

	// Подписка на топик
	mqtt.Subscribe("device/sensor", func(client mqtt.Client, msg mqtt.Message) {
		// Обрабатываем полученное сообщение
		deviceID := msg.Topic()
		data := string(msg.Payload()) // Для простоты принимаем строковые данные, их можно парсить

		// Для простоты считаем, что данные представляют собой число
		var value float64
		_, err := fmt.Sscanf(data, "%f", &value)
		if err != nil {
			log.Printf("Error parsing data: %v", err)
			return
		}

		// Создание записи для сохранения в базу данных
		deviceData := repository.DeviceData{
			DeviceID:  deviceID,
			Timestamp: time.Now().Format(time.RFC3339),
			DataType:  "sensor", // Тип данных, можно расширить
			Value:     value,
		}

		// Сохранение данных в базе
		err = db.SaveDeviceData(deviceData)
		if err != nil {
			log.Printf("Error saving data to database: %v", err)
		} else {
			log.Printf("Data saved: %+v", deviceData)
		}
	})

	// Ожидание сообщений
	select {}
}
