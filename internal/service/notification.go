package service

import (
	"practice/internal/domain"
	"practice/internal/port/in"
	"practice/internal/port/out"
)

var _ in.NotificationService = (*NotificationService)(nil)

type NotificationService struct {
	sender out.NotificationSender
}

func NewNotificationService(sender out.NotificationSender) *NotificationService {
	return &NotificationService{sender: sender}
}

func (n *NotificationService) Send(to string, message string) error {
	notification := domain.Notification{To: to, Message: message}
	if notification.IsEmpty() {
		return domain.ErrInvalidNotification
	}
	return n.sender.Send(to, message)
}
