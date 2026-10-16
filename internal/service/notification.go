package service

import (
	"practice/internal/port"
)

type NotificationService struct {
	sender port.Sender
}

func NewNotificationService(sender port.Sender) *NotificationService {
	return &NotificationService{sender: sender}
}

func (n *NotificationService) Send(to string, message string) error {
	return n.sender.Send(to, message)
}
