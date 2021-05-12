package handlers

import (
	"net/http"

	"github.com/haroundjudzman/golang-microservice/data"
)

// ListAll returns all current products in database
func (p *Products) ListAll(w http.ResponseWriter, r *http.Request) {
	p.l.Println("[DEBUG] List all records in database")

	// Fetch all products
	prodList := data.GetProducts()

	// Serialise to JSON
	err := prodList.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}
