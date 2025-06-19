package main

import (
	"fmt"
)

func main() {
	evenOrNotEven()
}

// 🧠 1. Четное или нечетное
func evenOrNotEven() {
	var number int
	fmt.Scanf("%d", &number)
	if number%2 == 0 {
		fmt.Println("Число четное")
	} else {
		fmt.Println("Число нечетное")
	}
}
