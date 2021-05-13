package data_test

import (
	"testing"

	"github.com/haroundjudzman/golang-microservice/data"
	"github.com/stretchr/testify/assert"
)

func TestMissingNameReturnsErr(t *testing.T) {
	b := data.Burger{
		Price: 10.0,
	}

	v := data.NewValidation()
	err := v.Validate(b)
	assert.Len(t, err, 1)
}

func TestMissingPriceReturnsErr(t *testing.T) {
	b := data.Burger{
		Name:  "Cheeseburger",
		Price: -1,
	}

	v := data.NewValidation()
	err := v.Validate(b)
	assert.Len(t, err, 1)
}

func TestValidBurgerDoesNotReturnErr(t *testing.T) {
	b := &data.Burger{
		Name:  "Cheeseburger",
		Price: 3.90,
	}

	v := data.NewValidation()
	err := v.Validate(b)
	assert.Len(t, err, 0)
}
