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
	"github.com/NeiderFajardo/pkg/database"
)

type InternalServer struct {
	Server  *http.Server
	Handler *http.ServeMux
	// config *Config
}

func NewServer() *InternalServer {
	mux := http.NewServeMux()
	stack := middlewares.CreateStack(
		middlewares.IsAuthed,
		middlewares.LogResponse,
	)
	return &InternalServer{
		Server: &http.Server{
			Addr:    ":8080",
			Handler: stack(mux),
		},
		Handler: mux,
	}
}

func RegisterHandlers(iserver *InternalServer, handlers map[string]http.Handler) {
	for path, handler := range handlers {
		iserver.Handler.Handle(path, handler)
	}

}

func Run(iserver *InternalServer, dbClient *database.MongoDatabase) {
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
	defer database.CloseConnection(dbClient)

	if err := iserver.Server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Failed to shutdown server: %v", err)
	}
}
