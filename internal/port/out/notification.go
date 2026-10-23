// Package out holds outbound (driven) ports, implemented by infra adapters.
package out

type NotificationSender interface {
	Send(to string, message string) error
}
