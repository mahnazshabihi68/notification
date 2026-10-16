package handler

import (
	"encoding/json"
	"net/http"
	"practice/internal/application"
	"practice/internal/domain"
)

type NotificationHandler struct {
	notificationApplication application.NotificationApplication
}

func NewNotificationHandler(notification application.NotificationApplication) *NotificationHandler {
	return &NotificationHandler{notificationApplication: notification}
}

func (h *NotificationHandler) SendEmail(w http.ResponseWriter, r *http.Request) {
	h.send(w, r, h.notificationApplication.Send)
}

func (h *NotificationHandler) send(w http.ResponseWriter, r *http.Request, deliver func(to, message string) error) {
	var req domain.SendRequest
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
