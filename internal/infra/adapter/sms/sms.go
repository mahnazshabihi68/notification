package sms

import "fmt"

type Sender struct {
}

func NewSender() *Sender {
	return &Sender{}
}

func (s *Sender) Send(to string, message string) error {
	fmt.Printf("Sending sms to = %s message = %s\n", to, message)
	return nil
}
