package handlers

import (
	"context"
	"net/http"

	"github.com/haroundjudzman/golang-microservice/data"
)

// MiddlewareBurgerValidate validates the burger in request and calls next handler.
func (b *Burgers) MiddlewareBurgerValidate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		burger := &data.Burger{}

		err := data.FromJSON(burger, r.Body)
		if err != nil {
			b.l.Println("[ERROR] serialising burger", err)

			w.WriteHeader(http.StatusBadRequest)
			data.ToJSON(&GenericError{Message: err.Error()}, w)
			return
		}

		// Validate burger
		errs := b.v.Validate(burger)
		if len(errs) != 0 {
			b.l.Println("[ERROR] validating burger", errs)
			w.WriteHeader(http.StatusUnprocessableEntity)
			data.ToJSON(&ValidationError{Messages: errs.Errors()}, w)
			return
		}
		// Add burger to the context
		ctx := context.WithValue(r.Context(), KeyBurger{}, burger)
		r = r.WithContext(ctx)

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
