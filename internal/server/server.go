package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/NeiderFajardo/internal/server/middlewares"
)

type InternalServer struct {
	Server  *http.Server
	Handler *http.ServeMux
	// config *Config
	middlewares []middlewares.IMiddleware
}

func NewServer(middlewares []middlewares.IMiddleware) *InternalServer {
	mux := http.NewServeMux()
	return &InternalServer{
		Server: &http.Server{
			Addr:    ":8080",
			Handler: mux,
		},
		Handler:     mux,
		middlewares: middlewares,
	}
}

func RegisterHandlers(iserver *InternalServer, handlers map[string]http.Handler) {
	for path, handler := range handlers {
		registerHandler(iserver, path, handler)
	}

}

func registerHandler(iserver *InternalServer, path string, handler http.Handler) {
	for _, middleware := range iserver.middlewares {
		handler = middleware.HandlerFunc(handler)
	}
	iserver.Handler.Handle(path, handler)
}

func Run(iserver *InternalServer) {
	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Printf("Server listening on %s", iserver.Server.Addr)
		err := iserver.Server.ListenAndServe()
		if err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Println("Server shutdown")
		} else {
			log.Fatal(err)
		}
	}()

	<-ctx.Done()
	log.Println("Shutting down server...")

	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownRelease()

	if err := iserver.Server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Failed to shutdown server: %v", err)
	}
}
