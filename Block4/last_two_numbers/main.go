package main

import (
	"fmt"
)

func main() {
	lastTwoNumbers(1234)

}

// 20. 🔚 Последние две цифры
func lastTwoNumbers(number int) {
	fmt.Println(number % 100)
}
