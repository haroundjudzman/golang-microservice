package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/haroundjudzman/golang-microservice/handlers"
)

func main() {

	logger := log.New(os.Stdout, "golang-microservice ", log.LstdFlags)

	productHandler := handlers.NewProducts(logger)

	serveMux := http.NewServeMux()
	serveMux.Handle("/", productHandler)

	server := &http.Server{
		Addr:         ":9090",
		Handler:      serveMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// Use goroutine so ListenAndServe won't block
	go func() {
		logger.Println("Starting server on port 9090")

		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Block because reading from channel won't happen until there is a message to be consumed
	sig := <-sigChan
	logger.Println("Received terminate signal, gracefully shutting down", sig)

	// Let current operation to complete with 30 seconds grace period
	timeoutContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	server.Shutdown(timeoutContext)
}
