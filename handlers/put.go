package handlers

import (
	"net/http"

	"github.com/haroundjudzman/golang-microservice/data"
)

// Update updates the product with given ID
func (p *Products) Update(w http.ResponseWriter, r *http.Request) {

	// Fetch product from context
	prod := r.Context().Value(KeyProduct{}).(*data.Product)

	p.l.Println("[DEBUG] Updating record id", prod.ID)

	err := data.UpdateProduct(prod)
	if err == data.ErrProductNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Product not found", http.StatusInternalServerError)
		return
	}
}
