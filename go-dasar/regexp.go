package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile("a([a-z])i")

	fmt.Println(regex.MatchString("ami"))
	fmt.Println(regex.MatchString("abi"))
	fmt.Println(regex.MatchString("aJi"))

	search := regex.FindAllString("ami amu ari", -1)
	fmt.Println(search)
}
