package main

import "fmt"

func main() {
	fmt.Println(double(5))
	fmt.Println(double(10))
	fmt.Println(double(25))
}

func double(n int) int {
	return n * 2
}
