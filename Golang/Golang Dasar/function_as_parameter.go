package main

import "fmt"

func hellospam(name string, filter func(string) string) {
	hellof := filter(name)
	fmt.Println("Hello", hellof)
}

func spam(name string) string {
	if name == "Babi"{
		return "xxx"
	}else {
		return name
	}
}

func main() {
	result := spam("Bayu")
	fmt.Println(result)
}