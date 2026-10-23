// Package in holds inbound (driving) ports, used by handlers.
package in

type NotificationService interface {
	Send(to string, message string) error
}
