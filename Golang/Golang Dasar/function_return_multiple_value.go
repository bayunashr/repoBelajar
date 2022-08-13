package main

import "fmt"

func hello()(string, string, string){
	return "Muhammad","Bayu","Nashrullah"
}

func main() {
	fn, nn, _ := hello()
	fmt.Println(fn,nn)
}