package main

import "fmt"

func hello(firstName string, nickName string, lastName string){
	fmt.Println("Hello", firstName, nickName, lastName)
}

func main() {
	name := "Bayu"
	hello("Muhammad", name, "Nashrullah")
}