package main

import (
	"net/http"

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

	r.POST("/notify", func(c *gin.Context) {
		var req NotifyRequest

		// request bind
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid request",
			})
			return
		}

		// validation check
		if req.UserID == "" || req.Message == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "user_id and message required",
			})
			return
		}

		println("Sending notification to:", req.UserID)
		println("Message:", req.Message)

		c.JSON(http.StatusOK, gin.H{
			"status": "notification sent",
		})
	})

	r.Run(":8081")
}
