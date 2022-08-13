package main

import "fmt"

type HasName interface{
	GetName() string
}

func hello(nama HasName){
	fmt.Println("Halo",nama.GetName())
}

type Person struct{
	Name string
}

func (person Person) GetName() string{
	return person.Name
}

type Animal struct{
	Name string
}

func (animal Animal) GetName() string{
	return animal.Name
}

func main() {
	bayu := Person{
		Name: "Bayu",
	}
	hello(bayu)

	cow := Animal{
		Name: "Cow",
	}
	hello(cow)
}