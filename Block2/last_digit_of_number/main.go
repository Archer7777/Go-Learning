package main

import "fmt"

func main() {
	lastDigitOfNumber(123)
}

//7. 🔢 Последняя цифра числа
func lastDigitOfNumber(number int) {
	fmt.Println(number % 10)
}
