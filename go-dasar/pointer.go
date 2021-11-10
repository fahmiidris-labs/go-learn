package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	var address1 Address = Address{"Kab. Bogor", "Jawa Barat", "Indonesia"}
	var address2 *Address = &address1
	var address3 *Address = &address1

	address2.City = "Bandung"

	*address2 = Address{"Malang", "Jawa Timur", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	var address4 *Address = new(Address)
	address4.Country = "Indonesia"
	fmt.Println(address4)

	var alamat = Address{
		City:     "Bekasi",
		Province: "Jawa Barat",
		Country:  "",
	}
	var alamatPointer *Address = &alamat

	ChangeCountryToIndonesia(&alamat)

	fmt.Println(alamatPointer)
}
