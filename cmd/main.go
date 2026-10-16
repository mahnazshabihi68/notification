package cmd

import (
	"log"
	"net/http"
	"practice/api"
	"practice/internal/handler"
	"practice/internal/infra/adapter/email"
	"practice/internal/service"
)

func main() {
	srv := service.NewNotificationService(email.NewSender())
	h := handler.NewNotificationHandler(srv)
	routes := api.NewRoute(h)

	mux := http.NewServeMux()
	routes.Register(mux)
	log.Fatal(http.ListenAndServe("8080", mux))
}
