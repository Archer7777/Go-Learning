package main

import "fmt"

func main() {
	remainderOfDivision(17, 5)
	lastDigitOfNumber(123)
	meanArithmetic(5, 7, 10)
	ageInDays(20)
	squareArea(4)
}

//6. 🧾 Остаток от деления
func remainderOfDivision(a int, b int) {
	fmt.Println(a % b)
}

//7. 🔢 Последняя цифра числа
func lastDigitOfNumber(number int) {
	fmt.Println(number % 10)
}

//8. 💯 Среднее арифметическое
func meanArithmetic(a, b, c float64) {
	fmt.Println((a + b + c) / 3)
}

//9. 🎂 Возраст в днях
func ageInDays(age int) {
	fmt.Println(age * 365)
}

//10. 💡 Площадь квадрата
func squareArea(side int) {
	fmt.Println(side * side)
}
