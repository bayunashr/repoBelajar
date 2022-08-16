package main

import "fmt"

func endApp(){
	msgpanic := recover()
	if msgpanic == nil{
		fmt.Println("exit")
	}else {
		fmt.Println(msgpanic)	
	}
}

func runApp(pan bool){
	defer endApp()
	if pan == false {
		fmt.Println("running")
	}else {
		panic("pan = true")
	}
}

func main() {
	runApp(true)
	fmt.Println("Cek apakah masih jalan")
}