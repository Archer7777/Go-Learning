package main

import (
	"fmt"
)

func main() {
	maxTwoNumbers()
}

// 🧮 4. Максимум из двух чисел
func maxTwoNumbers() {
	var number1, number2 int
	fmt.Scanf("%d %d", &number1, &number2)
	if number1 > number2 {
		fmt.Println(number1)
	} else if number1 == number2 {
		fmt.Println("Числа равны")
	} else {
		fmt.Println(number2)
	}
}
