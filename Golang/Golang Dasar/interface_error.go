package main

import (
	"errors"
	"fmt"
)

func pembagian(i int, j int) (int,error){
	if j == 0{
		return 0, errors.New("Pembagi tidak boleh 0")
	}else {
		hasil := i/j
		return hasil, nil
	}
}

func main() {
	hasil, err := pembagian(100,10)
	fmt.Println(hasil)
	fmt.Println(err)

	if err != nil{
		fmt.Println(err)
	}else {
		fmt.Println(hasil)
	}
}