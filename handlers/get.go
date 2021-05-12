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
	err := data.ToJSON(prodList, w)
	if err != nil {
		p.l.Println("[ERROR] serialising product", err)
	}
}

// ListSingle returns one product from given id param
func (p *Products) ListSingle(w http.ResponseWriter, r *http.Request) {
	// Get id from URL
	id := getProductID(r)

	p.l.Println("[DEBUG] Get product ID", id)

	prod, err := data.GetProductByID(id)

	switch err {
	case nil:

	case data.ErrProductNotFound:
		p.l.Println("[ERROR] No matching product", err)
		w.WriteHeader(http.StatusNotFound)
		return
	default:
		p.l.Println("[ERROR] No matching product", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = data.ToJSON(prod, w)
	if err != nil {
		p.l.Println("[ERROR] serialising product", err)
	}
}
