package application

import (
	"net/http"

	"practice/internal/api"
	"practice/internal/handler"
	"practice/internal/infra/adapter/email"
	"practice/internal/service"
)

func New() http.Handler {
	sender := email.NewSender()
	notificationService := service.NewNotificationService(sender)
	notificationHandler := handler.NewNotificationHandler(notificationService)
	routes := api.NewRoute(notificationHandler)

	mux := http.NewServeMux()
	routes.Register(mux)
	return mux
}
