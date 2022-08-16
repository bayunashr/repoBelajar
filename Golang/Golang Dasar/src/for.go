package main

import "fmt"

func main() {
	
	for counter := 1; counter <=5; counter++{
		fmt.Println(counter)
	}

	arr := []string{"Halo","Saya","Bayu"}
	for i, j := range arr{
		fmt.Println(i," ",j)
	}
	for _, j := range arr{
		fmt.Println(j)
	}
}