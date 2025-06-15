package main

import "fmt"

func main() {
	costCalculation(50, 3)
}

//4. 💸 Расчёт стоимости
func costCalculation(price float64, count int) {
	fullPrice := price * float64(count)
	fmt.Println(fullPrice)
}
