package main

import "fmt"

func main() {
	counter := 0
	name := "Bayu"
	increment := func ()  {
		name := "Budi"
		counter++
		fmt.Println(counter, name)
	}
	increment()
	increment()
	increment()
	fmt.Println(name)
}