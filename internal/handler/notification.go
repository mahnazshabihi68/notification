package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"practice/internal/domain"
	"practice/internal/handler/request"
	"practice/internal/port/in"
)

type NotificationHandler struct {
	emailService in.NotificationService
}

func NewNotificationHandler(emailService in.NotificationService) *NotificationHandler {
	return &NotificationHandler{emailService: emailService}
}

func (h *NotificationHandler) SendEmail(w http.ResponseWriter, r *http.Request) {
	h.send(w, r, h.emailService)
}

func (h *NotificationHandler) send(w http.ResponseWriter, r *http.Request, svc in.NotificationService) {
	var req request.SendNotification
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	err := svc.Send(req.To, req.Message)
	if errors.Is(err, domain.ErrInvalidNotification) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(w, "failed to send notification", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
