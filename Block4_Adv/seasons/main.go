package main

import (
	"fmt"
)

func main() {
	seasons()
}

// 📅 3. Определение сезона
func seasons() {
	var monthNumber int
	fmt.Scanf("%d", &monthNumber)
	switch monthNumber {
	case 12, 1, 2:
		fmt.Println("Зима")
	case 3, 4, 5:
		fmt.Println("Весна")
	case 6, 7, 8:
		fmt.Println("Лето")
	case 9, 10, 11:
		fmt.Println("Осень")
	}
}
