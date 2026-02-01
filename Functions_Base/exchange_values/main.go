package main

import "fmt"

func main() {
	fmt.Println(swap(3, 5))
}

func swap(a, b int) (int, int) {
	return b, a
}
