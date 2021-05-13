package data

import (
	"fmt"
	"time"
)

// ErrBurgerNotFound raised when burger cannot be found in database.
var ErrBurgerNotFound = fmt.Errorf("Burger not found")

// Burger defines structure for API data
// The burgers are from Bob's Burgers
type Burger struct {
	ID        int     `json:"id"`
	Name      string  `json:"name" validate:"required"`
	Season    int     `json:"season"`
	Episode   int     `json:"episode"`
	Price     float32 `json:"price" validate:"required,gt=0"`
	CreatedOn string  `json:"-"`
	UpdatedOn string  `json:"-"`
	DeletedOn string  `json:"-"`
}

// Burgers is collection of Burger
type Burgers []*Burger

// GetBurgers returns all burger from database
func GetBurgers() Burgers {
	return burgerList
}

// GetBurgerByID returns single burger from given id.
// Returns ErrBurgerNotFound if no match is found.
func GetBurgerByID(id int) (*Burger, error) {
	i := findIndexByBurgerID(id)
	if i == -1 {
		return nil, ErrBurgerNotFound
	}

	return burgerList[i], nil
}

// AddBurger appends new burger with incremented ID
func AddBurger(b *Burger) {
	b.ID = getNextID()
	burgerList = append(burgerList, b)
}

// UpdateBurger updates existing burger with a given ID.
// It returns an error if burger is not found.
func UpdateBurger(b *Burger) error {
	i := findIndexByBurgerID(b.ID)
	if i == -1 {
		return ErrBurgerNotFound
	}

	burgerList[i] = b

	return nil
}

// Delete existing burger with a given ID.
// Returns an error if burger is not found.
func DeleteBurger(id int) error {
	i := findIndexByBurgerID(id)
	if i == -1 {
		return ErrBurgerNotFound
	}

	// Swap target burger with last burger in the list
	// and zero the new last burger before deleting
	burgerList[len(burgerList)-1], burgerList[i] = nil, burgerList[len(burgerList)-1]
	burgerList = burgerList[:len(burgerList)-1]

	return nil
}

// findIndexByBurgerID finds the index of burger
// in the database. Returns -1 when no match is found.
func findIndexByBurgerID(id int) int {
	for i, b := range burgerList {
		if b.ID == id {
			return i
		}
	}

	return -1
}

func getNextID() int {
	lastId := burgerList[len(burgerList)-1].ID
	return lastId + 1
}

var burgerList = []*Burger{
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
