package main

import "fmt"

func main() {

	name := "Bayu"
	
	switch name{
	case "Bayu" :
		fmt.Println("Bayu")
	case "Joko" :
		fmt.Println("Joko")
	default :
		fmt.Println("Default")
	}
	
	switch length := len(name); {
	case length>0 && length<6:
		fmt.Println("Lebih dari nol")
	case length>5 && length<11:
		fmt.Println("Lebih dari 5")
	default :
		fmt.Println("default")
	}
}