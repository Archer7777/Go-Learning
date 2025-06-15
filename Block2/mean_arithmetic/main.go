package main

import "fmt"

func main() {
	meanArithmetic(5, 7, 10)
}

//8. 💯 Среднее арифметическое
func meanArithmetic(a, b, c float64) {
	fmt.Println((a + b + c) / 3)
}
