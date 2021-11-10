package main

import "fmt"

func main() {
	name := "Fahmi Idris"
	counter := 0

	increment := func ()  {
		name := "Dadut"
		fmt.Println("Increment")
		counter++
		fmt.Println(name)
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}