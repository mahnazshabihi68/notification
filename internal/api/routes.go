package api

import (
	"net/http"
	"practice/internal/handler"
)

type Route struct {
	notificationHandler *handler.NotificationHandler
}

func NewRoute(notification *handler.NotificationHandler) *Route {
	return &Route{notificationHandler: notification}
}

func (r *Route) Register(mux *http.ServeMux) {
	mux.HandleFunc("POST /notification/email", r.notificationHandler.SendEmail)
}
