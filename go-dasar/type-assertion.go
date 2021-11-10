package main

import "fmt"

func random() interface{} {
	return "Ups"
}

func main() {
	var result interface{} = random()
	var resultString string = result.(string)

	fmt.Println(resultString)

	switch value := result.(type) {
	case string:
		fmt.Println("Value", value)
	case int:
		fmt.Println("Value", value)
	default:
		fmt.Println("Unknow type")
	}
}