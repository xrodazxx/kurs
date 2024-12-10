package mqtt

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// SubscribeToTopics подписывается на несколько топиков
func SubscribeToTopics(client mqtt.Client, topics []string) error {
	for _, topic := range topics {
		if token := client.Subscribe(topic, 1, messageHandler); token.Wait() && token.Error() != nil {
			return token.Error()
		}
		fmt.Printf("Подписан на топик: %s\n", topic)
	}
	return nil
}

// messageHandler обрабатывает полученные сообщения
func messageHandler(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Получено сообщение из топика %s: %s\n", msg.Topic(), msg.Payload())
}
