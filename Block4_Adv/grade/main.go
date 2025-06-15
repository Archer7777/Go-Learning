package main

import (
	"fmt"
)

func main() {
	grade()
}

// 📊 6. Оценка по баллам
func grade() {
	var number int
	fmt.Scanf("%d", &number)
	if number >= 90 && number <= 100 {
		fmt.Println("Отлично")
	} else if number >= 70 && number < 90 {
		fmt.Println("Хорошо")
	} else if number >= 50 && number < 70 {
		fmt.Println("Удовлетворительно")
	} else if number < 50 {
		fmt.Println("Неудовлетворительно")
	} else {
		fmt.Println("Нет такой оценки")
	}
}
