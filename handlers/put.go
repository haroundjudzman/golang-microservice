package handlers

import (
	"net/http"

	"github.com/haroundjudzman/golang-microservice/data"
)

// swagger:route PUT /burgers burgers editBurger
// Edit existing burger
//
// Responses:
//  204: noContentResponse
// 	400: badRequestResponse
// 	404: notFoundResponse
//	422: validationErrorResponse

// Update updates the burger with given ID
func (b *Burgers) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

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
