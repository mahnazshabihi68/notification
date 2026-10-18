package domain

type Notification struct {
	To      string
	Message string
}

func (n Notification) IsEmpty() bool {
	return n.To == "" || n.Message == ""
}
