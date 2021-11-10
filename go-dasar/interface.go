package main

import "fmt"

type HasName interface{
	GetName() string
}

type Person struct {
	Name string
}

type Animal struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func (animal Animal) GetName() string {
	return animal.Name
}

func sayHello(hasName HasName)  {
	fmt.Println("Hello", hasName.GetName())
}

func main() {
	var fahmi Person
	fahmi.Name = "Fahmi Idris"

	sayHello(fahmi)

	cat := Animal{
		Name: "Empus",
	}

	sayHello(cat)
}