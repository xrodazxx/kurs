package main

import (
	"Kurs/pkg/handler"
	"Kurs/pkg/listener"
	"Kurs/pkg/repository"
	"Kurs/pkg/server"
	"Kurs/pkg/service"
	topichandler "Kurs/pkg/topic_handler"
	"context"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"
)

func main() {
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "db",
		Port:     "5432",
		Username: "postgres",
		Password: "sosi",
		DBName:   "postgres",
		SSLMode:  "disable",
	})
	if err != nil {
		panic(err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	router := handlers.InitRoutes()
	serv := server.New(router)

	topichandlers := topichandler.NewTopicHandlers(services)
	lstnr := listener.NewMQTT("tcp://broker.emqx.io:1883", "data_subscriber")

	go func() {
		serv.Run()
	}()
	go func() {
		lstnr.Start(topichandlers.IDevice)
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop
	serv.Stop(context.Background())
	lstnr.Stop()
}
