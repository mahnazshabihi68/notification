package email

import "fmt"

type Sender struct {
}

func NewSender() *Sender {
	return &Sender{}
}

func (s *Sender) Send(to string, message string) error {
	fmt.Println("Sending email to = %s message = %s "+to, message)
	return nil
}
