package handlers

import (
	"net/http"

	"github.com/haroundjudzman/golang-microservice/data"
)

// Delete removes the product with given id.
func (p *Products) Delete(w http.ResponseWriter, r *http.Request) {
	id := getProductID(r)

	p.l.Println("[DEBUG] Deleting record id", id)

	err := data.DeleteProduct(id)
	if err == data.ErrProductNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Product not found", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
