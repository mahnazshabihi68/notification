package request

type SendNotification struct {
	To      string `json:"to"`
	Message string `json:"message"`
}
