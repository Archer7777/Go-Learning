package main

import "fmt"

func main() {
	timeConverter(135)
}

//5. 🕐 Конвертер времени
func timeConverter(minutes int) {
	fmt.Printf("Вывод: %d часа %d минут!\n", (minutes / 60), (minutes % 60))
}
