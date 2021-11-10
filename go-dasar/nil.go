package main

import "fmt"

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	person := newMap("Fahmi Idris")
	if person == nil {
		fmt.Println("Person Kosong")
	} else {
		fmt.Println(person)
	}
}