package main

import (
	"github.com/xrodazxx/App-ToDO/pkg/handler"
	"github.com/xrodazxx/App-ToDO/pkg/repository"
	"github.com/xrodazxx/App-ToDO/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
}
