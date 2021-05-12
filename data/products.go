package data

import (
	"fmt"
	"time"

	"github.com/go-playground/validator"
)

// ErrProductNotFound raised when product cannot be found in database.
var ErrProductNotFound = fmt.Errorf("Product not found")

// Product defines structure for API data
// The products are burgers from Bob's Burgers
type Product struct {
	ID        int     `json:"id"`
	Name      string  `json:"name" validate:"required"`
	Season    int     `json:"season"`
	Episode   int     `json:"episode"`
	Price     float32 `json:"price" validate:"gt=0"`
	CreatedOn string  `json:"-"`
	UpdatedOn string  `json:"-"`
	DeletedOn string  `json:"-"`
}

func (p *Product) Validate() error {
	validate := validator.New()
	return validate.Struct(p)
}

// Products is collection of Product
type Products []*Product

// GetProducts returns all product from database
func GetProducts() Products {
	return productList
}

// GetProductByID returns single product from given id.
// Returns ErrProductNotFound if no match is found.
func GetProductByID(id int) (*Product, error) {
	i := findIndexByProductID(id)
	if i == -1 {
		return nil, ErrProductNotFound
	}

	return productList[i], nil
}

// AddProduct appends new product with incremented ID
func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

// UpdateProduct updates existing product with a given ID
// It returns an error if product is not found
func UpdateProduct(id int, p *Product) error {
	_, index, err := findProduct(id)
	if err != nil {
		return err
	}

	p.ID = id
	productList[index] = p

	return nil
}

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}

	return nil, -1, ErrProductNotFound
}

func findIndexByProductID(id int) int {
	for i, p := range productList {
		if p.ID == id {
			return i
		}
	}

	return -1
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
