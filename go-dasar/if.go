package main

import "fmt"

func main() {
	var name = "Gatau"

	if name == "Fahmi" {
		fmt.Println("Hello Fahmi")
	} else if name == "Dadut" {
		fmt.Println("Hello Dadut")
	} else if name == "Budi" {
		fmt.Println("Hello Budi")
	} else {
		fmt.Println("Hi, kenalan donk")
	}

	if length := len(name); length > 5 {
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
