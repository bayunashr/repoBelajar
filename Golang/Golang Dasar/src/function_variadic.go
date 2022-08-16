package main

import "fmt"

func sigma(numbers...int) int{
	total := 0
	for _, j := range numbers{
		total += j
	}
	return total
}

func main() {
	total := sigma(1,2,3,4,5)
	fmt.Println(total)

	slice := []int{1,2,3,4,5}
	total = sigma(slice...)
	fmt.Println(total)
}