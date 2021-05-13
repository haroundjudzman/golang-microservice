package handlers

import (
	"net/http"

	"github.com/haroundjudzman/golang-microservice/data"
)

// Delete removes the burger with given id.
func (b *Burgers) Delete(w http.ResponseWriter, r *http.Request) {
	id := getBurgerID(r)

	b.l.Println("[DEBUG] Deleting burger with id", id)

	err := data.DeleteBurger(id)
	if err == data.ErrBurgerNotFound {
		b.l.Println("[ERROR] Burger id does not exist")
		w.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, w)
		return
	}

	if err != nil {
		b.l.Println("[ERROR] Deleting burger")
		w.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
