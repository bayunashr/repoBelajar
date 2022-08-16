package main

import "fmt"

func main() {
	var nama [2]string
	nama[0]="Bayu"
	nama[1]="Nashrullah"
	fmt.Println(nama[0])
	fmt.Println(nama[1])

	var nomor = [5]int{
		1,2,3,4,5,
	}
	fmt.Println(nomor)
	fmt.Println(nomor[0])
	fmt.Println(len(nomor))
}