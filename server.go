package kurs

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Server struct {
	httpServer *http.Server
	db         *sql.DB
	mqttClient mqtt.Client
}

// Запуск сервера
func (s *Server) Run(port string, handler http.Handler, dbDSN string, mqttBroker string) error {
	// базa данных
	var err error
	s.db, err = sql.Open("postgres", dbDSN)
	if err != nil {
		return err
	}

	// Инициализация MQTT
	s.mqttClient = s.initMQTT(mqttBroker)

	// Запуск сервераpackage service

	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return s.httpServer.ListenAndServe()
}

// Завершение работы
func (s *Server) Shutdown(ctx context.Context) error {
	if s.mqttClient != nil {
		s.mqttClient.Disconnect(250)
	}
	if s.db != nil {
		s.db.Close()
	}
	return s.httpServer.Shutdown(ctx)
}

// MQTT-клиент
func (s *Server) initMQTT(broker string) mqtt.Client {
	opts := mqtt.NewClientOptions().AddBroker(broker).SetClientID("iot-monitor")
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	return client
}
