package main

import "fmt"

func main() {
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke", counter)
	}

	slice := []string{"Fahmi", "Idris", "Dadut"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for i, value := range slice {
		fmt.Println("Index", i, "=", value)
	}

	person := map[string]string{
		"name":  "Fahmi",
		"title": "Pelajar",
	}

	for key, value := range person {
		fmt.Println("Key", key, "=", value)
	}
}
