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
	result := ""
	fmt.Println(hello(result))
}