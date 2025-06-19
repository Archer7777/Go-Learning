package main

import "fmt"

func main() {
	fmt.Println(square(5) + square(7))
}

func square(n int) int {
	return n * n
}
