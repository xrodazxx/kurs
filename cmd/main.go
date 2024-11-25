package main

import (
	"Kurs/pkg/handler"
	"Kurs/pkg/repository"
	"Kurs/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
}
