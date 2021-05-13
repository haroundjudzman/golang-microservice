package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/haroundjudzman/golang-microservice/data"
)

// Burgers is http.Handler
type Burgers struct {
	l *log.Logger
}

func NewBurgers(l *log.Logger) *Burgers {
	return &Burgers{l}
}

// KeyBurger is a key used for Burger object in the context.
type KeyBurger struct {
}

// MiddlewareBurgerValidate validates the burger in request and calls next handler.
func (b *Burgers) MiddlewareBurgerValidate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		prod := &data.Burger{}

		err := data.FromJSON(prod, r.Body)
		if err != nil {
			b.l.Println("[ERROR] serialising burger", err)

			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Validate burger
		err = prod.Validate()
		if err != nil {
			b.l.Println("[ERROR] validating burger", err)
			http.Error(
				w,
				fmt.Sprintf("Unable to validate burger: %s", err),
				http.StatusBadRequest,
			)
			return
		}
		// Add burger to the context
		ctx := context.WithValue(r.Context(), KeyBurger{}, prod)
		r = r.WithContext(ctx)

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

// getBurgerID returns burger ID from URL.
// It should never fail because router ensures
// that path will produce valid number.
// Panic in the extreme case it fails.
func getBurgerID(r *http.Request) int {

	// Parse id from request param
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}

	return id
}
