package main

import "fmt"

func main() {
	remainderOfDivision(17, 5)
	lastDigitOfNumber(123)
	meanArithmetic(5, 7, 10)
	ageInDays(20)
	squareArea(4)
}

func remainderOfDivision(a int, b int) {
	fmt.Println(a % b)
}

func lastDigitOfNumber(number int) {
	fmt.Println(number % 10)
}

func meanArithmetic(a, b, c float64) {
	fmt.Println((a + b + c) / 3)
}

func ageInDays(age int) {
	fmt.Println(age * 365)
}

func squareArea(side int) {
	fmt.Println(side * side)
}
