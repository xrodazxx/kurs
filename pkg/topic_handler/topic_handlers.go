package topichandler

import (
	"Kurs/pkg/service"
)

type TopicHandlers struct {
	IDevice
}

func NewTopicHandlers(services *service.Service) *TopicHandlers {
	return &TopicHandlers{
		IDevice: NewDevice(services.IDevice),
	}
}
