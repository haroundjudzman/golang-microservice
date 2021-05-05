package data

import "time"

// Product defines structure for API data
// The products are burgers from Bob's Burgers
type Product struct {
	ID        int
	Name      string
	Season    int
	Episode   int
	Price     float32
	CreatedOn string
	UpdatedOn string
	DeletedOn string
}

var productList = []*Product{
	{
		ID:        1,
		Name:      "New Bacon-ings Burger",
		Season:    1,
		Episode:   1,
		Price:     9.40,
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
	{
		ID:        2,
		Name:      "Egger Can't Be Cheesers Burger",
		Season:    3,
		Episode:   11,
		Price:     8.40,
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
}
