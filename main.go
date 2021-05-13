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
	"github.com/haroundjudzman/golang-microservice/data"
	"github.com/haroundjudzman/golang-microservice/handlers"
)

func main() {

	l := log.New(os.Stdout, "golang-microservice ", log.LstdFlags)
	v := data.NewValidation()

	burgerHandler := handlers.NewBurgers(l, v)

	// Create a router
	r := mux.NewRouter()

	// Divide each methods into its own subrouter
	getRouter := r.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/burgers", burgerHandler.ListAll)
	getRouter.HandleFunc("/burgers/{id:[0-9]+}", burgerHandler.ListSingle)

	putRouter := r.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/burgers", burgerHandler.Update)
	putRouter.Use(burgerHandler.MiddlewareBurgerValidate)

	postRouter := r.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/burgers", burgerHandler.Create)
	postRouter.Use(burgerHandler.MiddlewareBurgerValidate)

	deleteRouter := r.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/burgers/{id:[0-9]+}", burgerHandler.Delete)

	server := &http.Server{
		Addr:         ":9090",
		Handler:      r,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     l,
	}

	// Use goroutine so ListenAndServe won't block
	go func() {
		l.Println("Starting server on port 9090")

		err := server.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// Set up trap for interrupt and sigterm signal
	// Use this mechanism to gracefully shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Block because reading from channel won't happen until there is a message to be consumed
	sig := <-sigChan
	l.Println("Received terminate signal, gracefully shutting down", sig)

	// Let current operation to complete with 30 seconds grace period
	timeoutContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	server.Shutdown(timeoutContext)
}
