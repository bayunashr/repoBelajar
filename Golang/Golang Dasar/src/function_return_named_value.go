package main

import "fmt"

func hello()(firstName, nickName, lastName string, birth int){
	firstName = "Muhammad"
	nickName = "Bayu"
	lastName = "Nashrullah"
	birth = 7

	return
}

func main() {
	fn, nn, _, bd := hello()
	fmt.Println(fn,nn,bd)
}