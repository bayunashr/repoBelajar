package main

import "fmt"

func main() {
	var(
		a=-1
		b=2
		c=5
	)

	var i=a+b+c
	fmt.Println(i)

	i*=c // i = i * c
	fmt.Println(i)
}