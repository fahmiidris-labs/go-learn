package main

import "fmt"

func main() {
	name := "Fahmi"

	switch name {
	case "Fahmi":
		fmt.Println("Hello Fahmi")
		fmt.Println("Hello Fahmi")
	case "Dadut":
		fmt.Println("Hello Dadut")
		fmt.Println("Hello Dadut")
	default:
		fmt.Println("Kenalan Dong")
		fmt.Println("Kenalan Dong")
	}

	// switch length := len(name); length > 5 {
	// case true:
	// 	fmt.Println("Nama Terlalu Panjang")
	// case false:
	// 	fmt.Println("Nama Sudah Benar")
	// }

	length := len(name)

	switch {
	case length > 10:
		fmt.Println("Nama Terlalu Panjang")
	case length > 5:
		fmt.Println("Nama Lumayan Panjang")
	default:
		fmt.Println("Nama Sudah Benar")
	}

}
