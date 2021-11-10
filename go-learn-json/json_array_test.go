package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArray(t *testing.T) {
	customer := &Customer{
		FirstName:  "Fahmi",
		MiddleName: "idris",
		LastName:   "A",
		Age:        19,
		Married:    false,
		Hobbies: []string{
			"Ngegame",
			"Ngoding",
		},
	}

	bytes, _ := json.Marshal(customer)

	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Fahmi","MiddleName":"idris","LastName":"A","Age":19,"Married":false,"Hobbies":["Ngegame","Ngoding"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.MiddleName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Age)
	fmt.Println(customer.Married)
	fmt.Println(customer.Hobbies)
	fmt.Println(customer.Addresses)
}

func TestJSONArrayComplex(t *testing.T) {
	customer := &Customer{
		FirstName: "Fahmi Idris",
		Addresses: []Address{
			{
				Street:     "Gatau",
				Country:    "Indonesia",
				PostalCode: "123456",
			},
			{
				Street:     "Gatau",
				Country:    "Indonesia",
				PostalCode: "123456",
			},
		},
	}

	bytes, _ := json.Marshal(customer)

	fmt.Println(string(bytes))
}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"Fahmi Idris","MiddleName":"","LastName":"","Age":0,"Married":false,"Hobbies":null,"Adreess":[{"Street":"Gatau","Country":"Indonesia","PostalCode":"123456"},{"Street":"Gatau","Country":"Indonesia","PostalCode":"123456"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.MiddleName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Age)
	fmt.Println(customer.Married)
	fmt.Println(customer.Hobbies)
	fmt.Println(customer.Addresses)
}

func TestOnlyJSONArrayComplexDecode(t *testing.T) {
	jsonString := `[{"Street":"Gatau","Country":"Indonesia","PostalCode":"123456"},{"Street":"Gatau","Country":"Indonesia","PostalCode":"123456"}]`
	jsonBytes := []byte(jsonString)

	addresses := &[]Address{}

	err := json.Unmarshal(jsonBytes, addresses)
	if err != nil {
		panic(err)
	}

	fmt.Println(addresses)
}

func TestOnlyJSONArrayComplex(t *testing.T) {
	addresses := []Address{
		{
			Street:     "Gatau",
			Country:    "Indonesia",
			PostalCode: "123456",
		},
		{
			Street:     "Gatau",
			Country:    "Indonesia",
			PostalCode: "123456",
		},
	}

	bytes, _ := json.Marshal(addresses)

	fmt.Println(string(bytes))
}
