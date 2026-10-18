package handler

import (
	"encoding/json"
	"net/http"
	"practice/internal/dto"
	"practice/internal/port"
)

type NotificationHandler struct {
	notificationService port.NotificationService
}

func NewNotificationHandler(notification port.NotificationService) *NotificationHandler {
	return &NotificationHandler{notificationService: notification}
}

func (h *NotificationHandler) SendEmail(w http.ResponseWriter, r *http.Request) {
	h.send(w, r, h.notificationService.Send)
}

func (h *NotificationHandler) send(w http.ResponseWriter, r *http.Request, deliver func(to, message string) error) {
	var req dto.SendRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	if err := deliver(req.To, req.Message); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
