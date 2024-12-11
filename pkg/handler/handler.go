package handler

import (
	"Kurs/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.LoadHTMLGlob("template/*")
	// авторизация аунтефикация
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	// маршрутка для устройств
	api := router.Group("/api") //, h.userIdentity)
	{
		devices := api.Group("/devices")
		{
			devices.POST("/", h.addDevices)            // Добавление устройства
			devices.GET("/", h.getAllDevices)          // Получение Списка Устройств
			devices.GET("/:id/on", h.infoDevices)      // Получение Информации о устройстве
			devices.POST("/:id/on", h.turnOnDevices)   // Turn ON
			devices.POST("/:id/off", h.turnOffDevices) // Turn Off
		}
	}
	return router
}
