package main

import "fmt"

func main() {
	var nilai32 int32 = 100000
	var niali64 int64 = int64(nilai32)

	fmt.Println(nilai32)
	fmt.Println(niali64)

	var name = "Fahmi"
	var f = name[0]
	var fString string = string(f)
	fmt.Println(name)
	fmt.Println(fString)
}
