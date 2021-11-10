package main

import "fmt"

func getFullName() (firstName, middleName, lastName string) {
	firstName = "Fahmi"
	middleName = "Idris"
	lastName = "A"
	return
}

func main() {
	firstName, middleName, lastName := getFullName()
	fmt.Println(firstName, middleName, lastName)
}