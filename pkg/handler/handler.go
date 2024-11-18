package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	// авторизация аунтефикация
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up")
		auth.POST("/sign-in")
	}
	// маршрутка для устройств
	api := router.Group("/api")
	{
		devices := api.Group("/devices")
		{
			devices.POST("/")        // Добавление устройства
			devices.GET("/")         // Получение Списка Устройств
			devices.GET("/:id/on")   // Получение Информации о устройстве
			devices.POST("/:id/on")  // Turn ON
			devices.POST("/:id/off") // Turn Off
		}
	}
	return router
}
