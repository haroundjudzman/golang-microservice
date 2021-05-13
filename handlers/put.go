package handlers

import (
	"net/http"

	"github.com/haroundjudzman/golang-microservice/data"
)

// Update updates the burger with given ID
func (b *Burgers) Update(w http.ResponseWriter, r *http.Request) {

	// Fetch burger from context
	prod := r.Context().Value(KeyBurger{}).(*data.Burger)

	b.l.Println("[DEBUG] Updating record id", prod.ID)

	err := data.UpdateBurger(prod)
	if err == data.ErrBurgerNotFound {
		http.Error(w, "Burger not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Burger not found", http.StatusInternalServerError)
		return
	}
}
