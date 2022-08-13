package main

import "fmt"

func log()  {
	fmt.Println("log success")
}

func runApp(value int)  {
	defer log()
	result := 10/value
	fmt.Println(result)
}

func main() {
	runApp(0)
}