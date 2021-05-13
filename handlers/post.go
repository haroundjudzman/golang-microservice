package handlers

import (
	"net/http"

	"github.com/haroundjudzman/golang-microservice/data"
)

// Create adds a burger to database
func (b *Burgers) Create(w http.ResponseWriter, r *http.Request) {
	b.l.Println("[DEBUG] Creating burger")

	// Fetch burger from context
	burger := r.Context().Value(KeyBurger{}).(*data.Burger)

	data.AddBurger(burger)
}
