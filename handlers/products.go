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

// Products is http.Handler
type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// KeyProduct is a key used for Product object in the context.
type KeyProduct struct {
}

// MiddlewareProductValidate validates the product in request and calls next handler.
func (p *Products) MiddlewareProductValidate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		prod := &data.Product{}

		err := data.FromJSON(prod, r.Body)
		if err != nil {
			p.l.Println("[ERROR] serialising product", err)

			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Validate product
		err = prod.Validate()
		if err != nil {
			p.l.Println("[ERROR] validating product", err)
			http.Error(
				w,
				fmt.Sprintf("Unable to validate product: %s", err),
				http.StatusBadRequest,
			)
			return
		}
		// Add product to the context
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

// getProductID returns product ID from URL.
// It should never fail because router ensures
// that path will produce valid number.
// Panic in the extreme case it fails.
func getProductID(r *http.Request) int {

	// Parse id from request param
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}

	return id
}
