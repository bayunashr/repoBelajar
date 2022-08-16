package main

import "fmt"

func main() {
	var name1="Bayu"
	var name2="bayu"
	var result=name1==name2
	fmt.Println(result)

	var(
		x=5
		y=7
	)
	fmt.Println(x<y)
	fmt.Println(x>y)
}