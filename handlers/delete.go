package handlers

import (
	"net/http"

	"github.com/haroundjudzman/golang-microservice/data"
)

// Delete removes the burger with given id.
func (b *Burgers) Delete(w http.ResponseWriter, r *http.Request) {
	id := getBurgerID(r)

	b.l.Println("[DEBUG] Deleting record id", id)

	err := data.DeleteBurger(id)
	if err == data.ErrBurgerNotFound {
		http.Error(w, "Burger not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Burger not found", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
