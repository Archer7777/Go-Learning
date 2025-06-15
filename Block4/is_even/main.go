package main

import (
	"fmt"
)

func main() {
	isEven(9)

}

// 16. 🧠 Чётность через остаток
func isEven(number int) {
	fmt.Println(number % 2)
}
