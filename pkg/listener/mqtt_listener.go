package listener

import (
	topichandler "Kurs/pkg/topic_handler"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MQTT struct {
	client mqtt.Client
}

func NewMQTT(brokerURL, clientID string) *MQTT {
	opts := mqtt.NewClientOptions().AddBroker(brokerURL).SetClientID(clientID)
	client := mqtt.NewClient(opts)

	return &MQTT{client: client}
}

func (m *MQTT) Start(deviceHandler topichandler.IDevice) {
	if token := m.client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Ошибка подключения к брокеру: %v", token.Error())
	}

	topics := []string{"devices/temperature/1", "devices/humidity/2", "devices/socket/3", "devices/sensor/4"}
	for _, topic := range topics {
		if token := m.client.Subscribe(topic, 0, deviceHandler.Save); token.Wait() && token.Error() != nil {
			log.Printf("Ошибка подписки на топик %s: %v", topic, token.Error())
		} else {
			log.Printf("Успешно подписан на топик: %s", topic)
		}
	}
}

func (m *MQTT) Stop() {
	m.client.Disconnect(250)
	log.Println("MQTT клиент отключен от брокера")
}
