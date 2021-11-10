package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Fahmi Idris", "Fahmi"))
	fmt.Println(strings.Contains("Fahmi Idris", "Budi"))

	fmt.Println(strings.Split("Fahmi Indris A", " "))

	fmt.Println(strings.ToLower("Fahmi Idris A"))
	fmt.Println(strings.ToUpper("Fahmi Idris A"))
	fmt.Println(strings.ToTitle("Fahmi Idris A"))

	fmt.Println(strings.Trim("      Fahmi Idris     ", " "))
	fmt.Println(strings.ReplaceAll("Fahmi Joko Fahmi", "Fahmi", "Budi"))
}