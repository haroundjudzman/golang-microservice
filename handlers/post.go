package handlers

import (
	"net/http"

	"github.com/haroundjudzman/golang-microservice/data"
)

// swagger:route POST /burgers burgers createBurger
// Create a new burger and insert to database
//
// Responses:
// 	200: burgersResponse
// 	400: badRequestResponse
//	422: validationErrorResponse

// Create adds a burger to database
func (b *Burgers) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	// Fetch burger from context
	burger := r.Context().Value(KeyBurger{}).(*data.Burger)

	b.l.Println("[DEBUG] Creating burger")
	data.AddBurger(burger)
}
