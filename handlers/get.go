package handlers

import (
	"net/http"

	"github.com/haroundjudzman/golang-microservice/data"
)

// ListAll returns all current burgers in database
func (b *Burgers) ListAll(w http.ResponseWriter, r *http.Request) {
	b.l.Println("[DEBUG] List all burgers in database")

	// Fetch all burgers
	burgerList := data.GetBurgers()

	// Serialise to JSON
	err := data.ToJSON(burgerList, w)
	if err != nil {
		b.l.Println("[ERROR] serialising burger", err)
	}
}

// ListSingle returns one burger from given id param
func (b *Burgers) ListSingle(w http.ResponseWriter, r *http.Request) {
	// Get id from URL
	id := getBurgerID(r)

	b.l.Println("[DEBUG] Get burger ID", id)

	burger, err := data.GetBurgerByID(id)

	switch err {
	case nil:

	case data.ErrBurgerNotFound:
		b.l.Println("[ERROR] No matching burger", err)
		w.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, w)
		return
	default:
		b.l.Println("[ERROR] No matching burger", err)
		w.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, w)
		return
	}

	err = data.ToJSON(burger, w)
	if err != nil {
		b.l.Println("[ERROR] serialising burger", err)
	}
}
