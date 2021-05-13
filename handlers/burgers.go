package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/haroundjudzman/golang-microservice/data"
)

// KeyBurger is a key used for Burger object in the context.
type KeyBurger struct{}

// Burgers is http.Handler
type Burgers struct {
	l *log.Logger
	v *data.Validation
}

// NewBurgers returns new Burgers with given logger and validator
func NewBurgers(l *log.Logger, v *data.Validation) *Burgers {
	return &Burgers{l, v}
}

// ErrInvalidProductPath is an error raised when product path is not valid
var ErrInvalidProductPath = fmt.Errorf("invalid path, path should be /products/{id}")

// GenericError is generic error message returned by server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
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
