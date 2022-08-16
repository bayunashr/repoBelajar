package main

import (
	"fmt"
)

func random() interface{}{
	return "Bayu"
}

func random2() interface{}{
	return 12
}

func main() {
	result := random()
	resultStr := result.(string)
	fmt.Println(resultStr)

	result2 := random2()
	switch value := result2.(type){
	case string :
		fmt.Println(value,"is string")
	case int :
		fmt.Println(value,"is int")
		default :
		fmt.Println("unknown")
	}
}