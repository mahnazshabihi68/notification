package dto

type SendRequest struct {
	To      string `json:"to"`
	Message string `json:"message"`
}
