package application

import (
	"context"
	"errors"
	"net/http"
	"time"

	"practice/internal/api"
	"practice/internal/handler"
	"practice/internal/infra/adapter/email"
	"practice/internal/service"
)

type App struct {
	server *http.Server
}

func New(addr string) *App {
	emailService := service.NewNotificationService(email.NewSender())
	notificationHandler := handler.NewNotificationHandler(emailService)
	routes := api.NewRoute(notificationHandler)

	mux := http.NewServeMux()
	routes.Register(mux)

	return &App{
		server: &http.Server{
			Addr:              addr,
			Handler:           mux,
			ReadHeaderTimeout: 5 * time.Second,
		},
	}
}

// Run blocks until ctx is cancelled or the server fails.
func (a *App) Run(ctx context.Context) error {
	errCh := make(chan error, 1)
	go func() {
		if err := a.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- err
		}
		close(errCh)
	}()

	select {
	case err := <-errCh:
		return err
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		return a.server.Shutdown(shutdownCtx)
	}
}
