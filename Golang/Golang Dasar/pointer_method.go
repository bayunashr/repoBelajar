package main

import "fmt"

type Person struct {
	Name string
}

func (man *Person) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	var bayu Person = Person{"Bayu"}
	bayu.Married()
	fmt.Println(bayu)
}
