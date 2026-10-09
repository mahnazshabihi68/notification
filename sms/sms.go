package sms

import "fmt"

type Sender struct{}

func New() *Sender {
	return &Sender{}
}

func (s *Sender) Send(to, message string) error {
	fmt.Printf("[sms] to=%s message=%s\n", to, message)
	return nil
}
