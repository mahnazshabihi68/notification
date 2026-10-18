package email

import "fmt"

type Sender struct {
}

func NewSender() *Sender {
	return &Sender{}
}

func (s *Sender) Send(to string, message string) error {
	fmt.Printf("Sending email to = %s message = %s\n", to, message)
	return nil
}
