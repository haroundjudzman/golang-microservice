package main

import (
	"fmt"
	"testing"

	"github.com/haroundjudzman/golang-microservice/sdk/client"
	"github.com/haroundjudzman/golang-microservice/sdk/client/burgers"
)

func TestClient(t *testing.T) {
	cfg := client.DefaultTransportConfig().WithHost("localhost:9090")
	c := client.NewHTTPClientWithConfig(nil, cfg)
	params := burgers.NewListBurgersParams()
	burger, err := c.Burgers.ListBurgers(params)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(burger)
}
