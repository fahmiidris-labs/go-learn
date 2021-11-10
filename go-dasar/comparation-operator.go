package main

import "fmt"

func main() {
	var name1 = "Fahmi"
	var nama2 = "Fahmi"

	var result bool = name1 == nama2
	fmt.Println(result)

	var value1 = 100
	var value2 = 200

	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
}
