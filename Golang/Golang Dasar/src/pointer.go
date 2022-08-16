package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func change(x *Address) {
	x.City = "Ekek"
}

func main() {

	// address1 := Address{"Sidoarjo", "Jawa Timur", "Indonesia"}
	// address2 := address1

	var address1 Address = Address{"Sidoarjo", "Jawa Timur", "Indonesia"}
	var address2 *Address = &address1
	var address3 *Address = &Address{"Bali", "Jawa Barat", "Indonesia"}
	var address4 *Address = new(Address)
	var address5 Address = Address{"Batu", "Jawa Timur", "Indonesia"}

	address2.Country = "Balkan"

	*address2 = Address{"Surabaya", "Jawa Timur", "Indonesia"}

	change(&address5)

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)
	fmt.Println(address4)
	fmt.Println(address5)

}
