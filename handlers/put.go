package handlers

import (
	"net/http"

	"github.com/haroundjudzman/golang-microservice/data"
)

// Update updates the burger with given ID
func (b *Burgers) Update(w http.ResponseWriter, r *http.Request) {

	// Fetch burger from context
	burger := r.Context().Value(KeyBurger{}).(*data.Burger)

	b.l.Println("[DEBUG] Updating record id", burger.ID)

	err := data.UpdateBurger(burger)
	if err == data.ErrBurgerNotFound {
		b.l.Println("[ERROR] burger not found")
		w.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
