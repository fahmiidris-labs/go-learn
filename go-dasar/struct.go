package main

import "fmt"

type Customer struct {
	Name, Address string
	Age int
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My Name Is", customer.Name)
}

func main() {
	var fahmi Customer

	fahmi.Name = "Fahmi Idris"
	fahmi.Address = "Kab. Bogor"
	fahmi.Age = 19

	fahmi.sayHi("Dadut")

	// fmt.Println(fahmi)

	// dadut := Customer{
	// 	Name: "Dadut",
	// 	Address: "Kab. Bogor",
	// 	Age: 19,
	// }

	// fmt.Println(dadut)
}