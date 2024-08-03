package handler

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
	"vdo-be/pkg/middleware"
)

func RunServer(
	ctx context.Context,
	getEnv func(string) string,
) error {
	server := newServer()

	httpServer := &http.Server{
		Addr:    net.JoinHostPort("0.0.0.0", "8080"),
		Handler: server,
	}
	go func() {
		log.Printf("listening and serving on %s\n", httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Fprintf(os.Stderr, "error listening and serving: %s\n", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	wait := 60 * time.Second
	shutdownCtx, cancel := context.WithTimeout(ctx, wait)
	defer cancel()

	if err := httpServer.Shutdown(shutdownCtx); err != nil {
		log.Printf("error shutting down http server: %s\n", err)
		return err
	}

	log.Printf("shut down http server with timeout: %s\n", wait)

	return nil

}

func newServer() http.Handler {
	mux := http.NewServeMux()
	addRoutes(
		mux,
	)

	var handler http.Handler = mux
	handler = middleware.CORSMiddleware()(handler)

	return handler
}
