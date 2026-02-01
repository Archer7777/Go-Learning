package main

import (
	"fmt"
)

func main() {
	positiveNegativeOrZero()
}

// 🔢 5. Число — положительное, отрицательное или ноль
func positiveNegativeOrZero() {
	var number int
	fmt.Scanf("%d", &number)
	if number < 0 {
		fmt.Println("Число отрицательное")
	} else if number > 0 {
		fmt.Println("Число положительное")
	} else {
		fmt.Println("Число Ноль")
	}
}
