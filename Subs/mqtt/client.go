package mqtt

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// InitializeClient инициализирует MQTT-клиент
func InitializeClient(broker string, clientID string) mqtt.Client {
	options := mqtt.NewClientOptions()
	options.AddBroker(broker)
	options.SetClientID(clientID)
	options.OnConnect = func(client mqtt.Client) {
		fmt.Println("Успешное подключение к MQTT-брокеру!")
	}
	options.OnConnectionLost = func(client mqtt.Client, err error) {
		fmt.Printf("Соединение потеряно: %v\n", err)
	}
	return mqtt.NewClient(options)
}
