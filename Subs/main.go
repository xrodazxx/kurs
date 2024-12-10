package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DeviceData struct {
	DeviceID   string  `db:"device_id"`
	DeviceName string  `db:"device_name"`
	Timestamp  string  `db:"timestamp"`
	DataType   string  `db:"data_type"`
	Value      float64 `db:"value"`
}

var db *sqlx.DB

func saveDeviceData(data DeviceData) error {
	query := "INSERT INTO device_data (device_id, device_name, timestamp, data_type, value) VALUES ($1, $2, NOW(), $3, $4)"
	_, err := db.Exec(query, data.DeviceID, data.DeviceName, data.DataType, data.Value)
	return err
}

func main() {
	// Подключение к PostgreSQL
	var err error
	db, err = sqlx.Connect("postgres", "user=postgres password=mysecretpassword dbname=mydb sslmode=disable")
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	defer db.Close()

	// Обработчик MQTT
	messageHandler := func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("Пришло сообщение: %s в топик: %s\n", msg.Payload(), msg.Topic())

		// Парсим сообщение
		parts := strings.Split(msg.Topic(), "/")
		if len(parts) != 3 {
			log.Printf("Неправильный формат топика: %s", msg.Topic())
			return
		}

		deviceData := DeviceData{
			DeviceID:   parts[2],
			DeviceName: string(msg.Payload()),
			DataType:   parts[1],
		}

		if err := saveDeviceData(deviceData); err != nil {
			log.Printf("Ошибка сохранения данных: %v", err)
		} else {
			log.Printf("Данные устройства %s успешно сохранены", deviceData.DeviceName)
		}
	}

	// Настройка MQTT
	opts := mqtt.NewClientOptions().AddBroker("tcp://broker.emqx.io:1883").SetClientID("data_subscriber")
	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Ошибка подключения к брокеру: %v", token.Error())
	}
	defer client.Disconnect(250)

	topics := []string{"devices/temperature/1", "devices/humidity/2", "devices/socket/3", "devices/sensor/4"}
	for _, topic := range topics {
		client.Subscribe(topic, 0, messageHandler)
	}

	// Ожидание завершения
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan
	log.Println("Завершаем подписку...")
}
