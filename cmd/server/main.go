package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/thetrame22/project-wander-realm/internal/server"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	srv, err := server.New()
	if err != nil {
		log.Fatal("could not create server")
	}

	go func() {
		if err := srv.Start(); err != nil {
			log.Printf("server stopped: %v", err)
		}
	}()

	log.Println("Server is running. Press Ctrl+C to exit...")

	<-ctx.Done()
	log.Println("\nShutdown signal received...")

	err = srv.Stop()
	if err != nil {
		log.Fatalf("Server forced to shutdown: %s", err)
	}

	log.Println("Server exiting gracefully.")
}
