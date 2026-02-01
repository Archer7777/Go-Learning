package main

import "fmt"

var name string = "Petya"

func main() {
	fmt.Println(name)
	changeName()
}

func changeName() {
	name = "Alex"
	fmt.Println(name)
}
