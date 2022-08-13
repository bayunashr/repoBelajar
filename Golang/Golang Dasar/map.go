package main

import "fmt"

func main() {
	newmap:=make(map[string]string)
	newmap["nama"]="Bayu"
	newmap["alamat"]="Sidoarjo"
	newmap["umur"]="19"
	fmt.Println(newmap)
	
	delete(newmap, "umur")
	fmt.Println(newmap)
}