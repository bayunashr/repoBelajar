package main

import "fmt"

func main() {
	day:=[7]string{
		"Senin","Selasa","Rabu","Kamis","Jumat","Sabtu","Minggu",
	}
	fmt.Println(day)

	slice1:=day[3:]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	slice2:=append(slice1,"Ahad")
	fmt.Println(slice2)
	fmt.Println(day)

	newslice:=make([]string, len(slice1), cap(slice1))
	copy(newslice,slice1)
	fmt.Println(newslice)
	fmt.Println(len(newslice))
	fmt.Println(cap(newslice))

	thisarray:=[...]int{1,2,3,4}
	thisslice:=[]int{1,2,3,4}
	fmt.Println(thisarray)
	fmt.Println(thisslice)
}