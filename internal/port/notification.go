package port

type Sender interface {
	Send(to string, message string) error
}
