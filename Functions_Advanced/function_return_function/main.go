package main

import "fmt"

func main() {
	double := makeMultiplier(2)
	fmt.Println(double(5))
}

func makeMultiplier(factor int) func(int) int {
	return func(input int) int {
		return input * factor
	}
}
