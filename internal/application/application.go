package application

type NotificationApplication interface {
	Send(to string, message string) error
}
