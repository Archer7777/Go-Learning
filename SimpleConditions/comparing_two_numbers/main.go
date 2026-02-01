package main

import "fmt"

func main() {
	comparingTwoNumbers(10, 9)
}

//13. 🔃 Сравнение двух чисел
func comparingTwoNumbers(a, b int) {
	if b > a {
		fmt.Println("Второе больше")
		return
	}

	fmt.Println("Первое больше")
}
