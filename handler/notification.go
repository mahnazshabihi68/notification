package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"practice/service"
)

type NotificationService interface {
	SendEmail(to, message string) error
	SendSMS(to, message string) error
	List() []service.SentMessage
}

type NotificationHandler struct {
	svc NotificationService
}

func NewNotification(svc NotificationService) *NotificationHandler {
	return &NotificationHandler{svc: svc}
}

func (h *NotificationHandler) Register(mux *http.ServeMux) {
	mux.HandleFunc("POST /notifications/email", h.sendEmail)
	mux.HandleFunc("POST /notifications/sms", h.sendSMS)
	mux.HandleFunc("GET /notifications", h.list)
}

type sendRequest struct {
	To      string `json:"to"`
	Message string `json:"message"`
}

func (h *NotificationHandler) sendEmail(w http.ResponseWriter, r *http.Request) {
	h.send(w, r, h.svc.SendEmail)
}

func (h *NotificationHandler) sendSMS(w http.ResponseWriter, r *http.Request) {
	h.send(w, r, h.svc.SendSMS)
}

func (h *NotificationHandler) send(w http.ResponseWriter, r *http.Request, send func(to, message string) error) {
	var req sendRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, errors.New("invalid json body"))
		return
	}
	if req.To == "" || req.Message == "" {
		writeError(w, http.StatusBadRequest, errors.New("to and message are required"))
		return
	}
	if err := send(req.To, req.Message); err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}
	writeJSON(w, http.StatusAccepted, map[string]string{"status": "sent"})
}

func (h *NotificationHandler) list(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, h.svc.List())
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}

func writeError(w http.ResponseWriter, status int, err error) {
	writeJSON(w, status, map[string]string{"error": err.Error()})
}
