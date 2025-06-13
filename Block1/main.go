package main

import "fmt"

func main() {
	sum(2, 5)
	divide(10, 2)
	smToMeters(250)
	costCalculation(50, 3)
	timeConverter(135)
}

func sum(a int, b int) {
	fmt.Println(a + b)
}

func divide(a float64, b float64) {
	fmt.Println(a / b)
}

func smToMeters(sm float64) {
	fmt.Printf("Вывод: %g м\n", (sm / 100))
}

func costCalculation(price float64, count int) {
	fullPrice := price * float64(count)
	fmt.Println(fullPrice)
}

func timeConverter(minutes int) {
	fmt.Printf("Вывод: %d часа %d минут!\n", (minutes / 60), (minutes % 60))
}
