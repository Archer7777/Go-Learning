package main

import "fmt"

func main() {
	countdown(5)
}

func countdown(n int) {
	for i := n; i > 0; i-- {
		fmt.Println(i)
	}
}
