package topichandler

import (
	"Kurs/pkg/model"
	"Kurs/pkg/service"
	"context"
	"fmt"
	"log"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type IDevice interface {
	Save(client mqtt.Client, msg mqtt.Message)
}

type Device struct {
	deviceService service.IDevice
}

func NewDevice(
	deviceService service.IDevice,
) *Device {
	return &Device{
		deviceService: deviceService,
	}
}

func (th *Device) Save(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Пришло сообщение: %s в топик: %s\n", msg.Payload(), msg.Topic())

	// Парсим сообщение
	parts := strings.Split(msg.Topic(), "/")
	if len(parts) != 3 {
		log.Printf("Неправильный формат топика: %s", msg.Topic())
		return
	}

	deviceData := model.DeviceData{
		DeviceID:   parts[2],
		DeviceName: string(msg.Payload()),
		DataType:   parts[1],
	}

	err := th.deviceService.Save(context.Background(), deviceData) //косяк хз откуда взять в msg контекст
	if err != nil {
		log.Println(err.Error())
	}
}
