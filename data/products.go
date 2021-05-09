package data

import (
	"encoding/json"
	"io"
	"time"
)

// Product defines structure for API data
// The products are burgers from Bob's Burgers
type Product struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Season    int     `json:"season"`
	Episode   int     `json:"episode"`
	Price     float32 `json:"price"`
	CreatedOn string  `json:"-"`
	UpdatedOn string  `json:"-"`
	DeletedOn string  `json:"-"`
}

func (p *Product) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(p)
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(p)
}

func GetProducts() Products {
	return productList
}

func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

func getNextID() int {
	lastId := productList[len(productList)-1].ID
	return lastId + 1
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
