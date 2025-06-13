package main

import "fmt"

func main() {
	sum(2, 5)
	divide(10, 2)
	smToMeters(250)
	costCalculation(50, 3)
	timeConverter(135)
}

//1. 🧮 Сложение двух чисел
func sum(a int, b int) {
	fmt.Println(a + b)
}

//2. ➗ Деление на два
func divide(a float64, b float64) {
	fmt.Println(a / b)
}

//3. 📏 Перевод сантиметров в метры
func smToMeters(sm float64) {
	fmt.Printf("Вывод: %g м\n", (sm / 100))
}

//4. 💸 Расчёт стоимости
func costCalculation(price float64, count int) {
	fullPrice := price * float64(count)
	fmt.Println(fullPrice)
}

//5. 🕐 Конвертер времени
func timeConverter(minutes int) {
	fmt.Printf("Вывод: %d часа %d минут!\n", (minutes / 60), (minutes % 60))
}
