package main

import "fmt"

type Blacklist func(string) bool

func regUser(name string, blacklist Blacklist){
	if blacklist(name) {
		fmt.Println("Blacklisted", name)
	}else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	result := func(name string) bool{
		return name == "Bayu"
	}

	regUser("admin", result)
}