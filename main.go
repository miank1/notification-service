package main

import (
	"net/http"
	"notification-service/handler"
	"notification-service/service"

	"github.com/gin-gonic/gin"
)

type NotifyRequest struct {
	UserID  string `json:"user_id"`
	Message string `json:"message"`
}

func main() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "notification service running",
		})
	})

	notificationService := service.NewNotificationService()
	notificationHandler := handler.NewNotificationHandler(notificationService)

	// routes
	r.POST("/notify", notificationHandler.SendNotification)

	r.Run(":8081")
}
