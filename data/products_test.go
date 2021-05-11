package data_test

import (
	"testing"

	"github.com/haroundjudzman/golang-microservice/data"
)

func TestChecksValidation(t *testing.T) {
	p := &data.Product{
		Name:  "Cheeseburger",
		Price: 3.90,
	}

	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
