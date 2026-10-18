package port

type NotificationService interface {
	Send(to string, message string) error
}

type SenderService interface {
	Send(to string, message string) error
}
