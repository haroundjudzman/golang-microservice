package data_test

import (
	"testing"

	"github.com/haroundjudzman/golang-microservice/data"
)

func TestChecksValidation(t *testing.T) {
	b := &data.Burger{
		Name:  "Cheeseburger",
		Price: 3.90,
	}

	err := b.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
