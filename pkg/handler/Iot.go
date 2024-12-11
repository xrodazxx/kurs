package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// добавление устройства
func (h *Handler) addDevices(c *gin.Context) {
}

// получение устройств
func (h *Handler) getAllDevices(c *gin.Context) {
	devices, err := h.services.IDevice.Get(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch devices"})
	}

	c.HTML(http.StatusOK, "devices.html", devices)
}

// информация о устройстве
func (h *Handler) infoDevices(c *gin.Context) {

}

// включение устройства
func (h *Handler) turnOnDevices(c *gin.Context) {

}

// выключение устройства
func (h *Handler) turnOffDevices(c *gin.Context) {

}
