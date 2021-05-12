package handlers

import (
	"net/http"

	"github.com/haroundjudzman/golang-microservice/data"
)

// Create adds a product to database
func (p *Products) Create(w http.ResponseWriter, r *http.Request) {
	p.l.Println("[DEBUG] Creating product")

	// Fetch product from context
	product := r.Context().Value(KeyProduct{}).(*data.Product)

	data.AddProduct(product)
}
