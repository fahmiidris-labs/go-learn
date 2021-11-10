package main

import "fmt"

func getFullName() (string, string, string) {
	return "Fahmi", "Idris", "A"
}

func main() {
	firstName, lastName, _  := getFullName()
	fmt.Println(firstName, lastName)
}