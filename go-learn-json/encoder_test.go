package main

import (
	"encoding/json"
	"os"
	"testing"
)

func TestEncoder(t *testing.T) {
	writer, _ := os.Create("CustomerOut.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName:  "Fahmi",
		MiddleName: "Idris",
		LastName:   "A",
	}

	encoder.Encode(customer)
}
