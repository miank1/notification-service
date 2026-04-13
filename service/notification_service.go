package service

import "notification-service/model"

type NotificationService struct{}

func NewNotificationService() *NotificationService {
	return &NotificationService{}
}

func (s *NotificationService) SendNotification(req model.NotifyRequest) error {
	// business logic (for now just print)
	println("Sending notification to:", req.UserID)
	println("Message:", req.Message)

	return nil
}
