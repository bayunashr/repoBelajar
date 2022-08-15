package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {

	// address1 := Address{"Sidoarjo", "Jawa Timur", "Indonesia"}
	// address2 := address1

	var address1 Address = Address{"Sidoarjo", "Jawa Timur", "Indonesia"}
	var address2 *Address = &address1
	var address3 *Address = &Address{"Bali", "Jawa Barat", "Indonesia"}
	var address4 *Address = new(Address)

	address2.Country = "Balkan"

	*address2 = Address{"Surabaya", "Jawa Timur", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)
	fmt.Println(address4)

}
