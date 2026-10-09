package main

import (
	"log"
	"net/http"

	"practice/email"
	"practice/handler"
	"practice/service"
	"practice/sms"
)

func main() {
	svc := service.NewNotification(email.New(), sms.New())
	h := handler.NewNotification(svc)

	mux := http.NewServeMux()
	h.Register(mux)

	log.Println("listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
