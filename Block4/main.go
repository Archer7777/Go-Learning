package main

import (
	"fmt"
	"math"
)

func main() {
	isEven(9)
	unsignedDifference(10, 17)
	tempInFarenheit(32)
	getInHand(100000)
	lastTwoNumbers(1234)

}

// 16. 🧠 Чётность через остаток
func isEven(number int) {
	fmt.Println(number % 2)
}

// 17. 🧮 Разность без знака
func unsignedDifference(a, b float64) {
	module := math.Abs(a - b)
	fmt.Println(module)
}

// 18. 🌡️ Температура в Фаренгейтах
func tempInFarenheit(celcius float64) {
	fmt.Println(celcius*1.8 + 32)
}

// 19. 💵 Сколько получишь на руки
func getInHand(salary float64) {
	fmt.Println(salary - salary*0.13)
}

// 20. 🔚 Последние две цифры
func lastTwoNumbers(number int) {
	fmt.Println(number % 100)
}
