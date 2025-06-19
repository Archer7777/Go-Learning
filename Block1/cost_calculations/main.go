package main

func main() {
	costCalculation(50, 3)
}

//4. 💸 Расчёт стоимости
func costCalculation(price float64, count int) float64 {
	fullPrice := price * float64(count)
	return fullPrice
}
