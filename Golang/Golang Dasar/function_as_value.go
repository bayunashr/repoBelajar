package main

import "fmt"

func hello(name string) string {
	if name == ""{
		return "Nama anda kosong"
	}else{
		return name
	}
}

func main() {
	vhello := hello
	result := vhello("Bayu")
	fmt.Println(result)
}