package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/haroundjudzman/golang-microservice/handlers"
)

func main() {

	logger := log.New(os.Stdout, "golang-microservice ", log.LstdFlags)

	productHandler := handlers.NewProducts(logger)

	// Create a router
	r := mux.NewRouter()

	// Divide each methods into its own subrouter
	getRouter := r.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/products", productHandler.ListAll)
	getRouter.HandleFunc("/products/{id:[0-9]+}", productHandler.ListSingle)

	putRouter := r.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/products", productHandler.Update)
	putRouter.Use(productHandler.MiddlewareProductValidate)

	postRouter := r.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/products", productHandler.Create)
	postRouter.Use(productHandler.MiddlewareProductValidate)

	server := &http.Server{
		Addr:         ":9090",
		Handler:      r,
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

	// Set up trap for interrupt and sigterm signal
	// Use this mechanism to gracefully shutdown
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
