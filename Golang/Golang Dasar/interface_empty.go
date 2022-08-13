package main

import (
	"fmt"
)

func test(i int) interface{}{
	if i == 1{
		return 1
	}else if i == 2{
		return false
	}else {
		return "string"
	}
}

func testAgain(j interface{}) interface{}{
	if j == true {
		return true
	}else if j == false{
		return false
	}else if j == "string"{
		return "string"
	}else{
		return 1
	}
}

func main() {
	var test interface{} = test(2)
	fmt.Println(test)

	var testagain interface{} = testAgain("halo")
	fmt.Println(testagain)
}