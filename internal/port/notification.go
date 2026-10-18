package port

type NotificationService interface {
	Send(to string, message string) error
}

type NotificationRepository interface {
	Send(to string, message string) error
}
