package handlers

import (
	"net/http"

	"github.com/haroundjudzman/golang-microservice/data"
)

// ListAll returns all current burgers in database
func (b *Burgers) ListAll(w http.ResponseWriter, r *http.Request) {
	b.l.Println("[DEBUG] List all records in database")

	// Fetch all burgers
	prodList := data.GetBurgers()

	// Serialise to JSON
	err := data.ToJSON(prodList, w)
	if err != nil {
		b.l.Println("[ERROR] serialising burger", err)
	}
}

// ListSingle returns one burger from given id param
func (b *Burgers) ListSingle(w http.ResponseWriter, r *http.Request) {
	// Get id from URL
	id := getBurgerID(r)

	b.l.Println("[DEBUG] Get burger ID", id)

	prod, err := data.GetBurgerByID(id)

	switch err {
	case nil:

	case data.ErrBurgerNotFound:
		b.l.Println("[ERROR] No matching burger", err)
		w.WriteHeader(http.StatusNotFound)
		return
	default:
		b.l.Println("[ERROR] No matching burger", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = data.ToJSON(prod, w)
	if err != nil {
		b.l.Println("[ERROR] serialising burger", err)
	}
}
