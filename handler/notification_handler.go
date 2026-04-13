package handler

import (
	"net/http"

	"notification-service/model"
	"notification-service/service"

	"github.com/gin-gonic/gin"
)

type NotificationHandler struct {
	service *service.NotificationService
}

func NewNotificationHandler(s *service.NotificationService) *NotificationHandler {
	return &NotificationHandler{service: s}
}

func (h *NotificationHandler) SendNotification(c *gin.Context) {
	var req model.NotifyRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}

	if req.UserID == "" || req.Message == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "user_id and message required",
		})
		return
	}

	err := h.service.SendNotification(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to send notification",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "notification sent",
	})
}
