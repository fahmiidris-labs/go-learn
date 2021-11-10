package main

import (
	"fmt"
	"go-dasar/database"
)

func main() {
	result :=  database.GetDatabase()
	fmt.Println(result)
}