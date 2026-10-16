package domain

type SendRequest struct {
	To      string `json:"to"`
	Message string `json:"message"`
}
