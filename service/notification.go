package service

type Sender interface {
	Send(to, message string) error
}

type SentMessage struct {
	Channel string `json:"channel"`
	To      string `json:"to"`
	Message string `json:"message"`
}

type Notification struct {
	email Sender
	sms   Sender
	sent  []SentMessage
}

func NewNotification(email, sms Sender) *Notification {
	return &Notification{
		email: email,
		sms:   sms,
		sent:  []SentMessage{},
	}
}

func (n *Notification) SendEmail(to, message string) error {
	if err := n.email.Send(to, message); err != nil {
		return err
	}
	n.sent = append(n.sent, SentMessage{Channel: "email", To: to, Message: message})
	return nil
}

func (n *Notification) SendSMS(to, message string) error {
	if err := n.sms.Send(to, message); err != nil {
		return err
	}
	n.sent = append(n.sent, SentMessage{Channel: "sms", To: to, Message: message})
	return nil
}

func (n *Notification) List() []SentMessage {
	return n.sent
}
