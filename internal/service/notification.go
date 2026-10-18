package service

import (
	"errors"
	"practice/internal/domain"
	"practice/internal/port"
)

type NotificationService struct {
	sender port.SenderService
}

func NewNotificationService(sender port.SenderService) *NotificationService {
	return &NotificationService{sender: sender}
}

func (n *NotificationService) Send(to string, message string) error {
	notification := domain.Notification{To: to, Message: message}
	if notification.IsEmpty() {
		return errors.New("to and message are required")
	}
	return n.sender.Send(to, message)
}
