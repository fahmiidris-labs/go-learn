package main

import "fmt"

type Blacklist func(string) bool

func resgisterUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are bloked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	Blacklist := func(name string) bool {
		return name == "Admin"
	}

	resgisterUser("Fahmi Idris", Blacklist)
	resgisterUser("Admin", Blacklist)
}