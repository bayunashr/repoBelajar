package main

import "fmt"

func factorialFor(value int)int{
	result := 1
	for i := value; i>=1; i--{
		result *= i
	}
	return result
}

func factorialRecursive(value int)int{
	result := 0
	if value == 1 {
		return 1
	}else {
		result = value * factorialRecursive(value-1)
		return result
	}
}

func main() {
	factfor := factorialFor(5)
	fmt.Println(factfor)
	factrec := factorialRecursive(10)
	fmt.Println(factrec)
}