package main

import "fmt"

func main() {
	equalToZero(1)
	numberMoreHundred(120)
	comparingTwoNumbers(10, 9)
	divideOnThree(10)
	ageVerification(7)
}

//11. 🔍 Равно ли нулю
func equalToZero(number int) {
	if number == 0 {
		fmt.Println("Число равно нулю")
		return
	}

	fmt.Println("Число не равно нулю")
}

//12. 🪞 Число больше 100?
func numberMoreHundred(number int) {
	if number > 100 {
		fmt.Println("Много")
	} else {
		fmt.Println("Норм")
	}
}

//13. 🔃 Сравнение двух чисел
func comparingTwoNumbers(a, b int) {
	if b > a {
		fmt.Println("Второе больше")
		return
	}

	fmt.Println("Первое больше")
}

//14. 📛 Делится на 3
func divideOnThree(number int) {
	if number%3 == 0 {
		fmt.Println("Делится")
		return
	}

	fmt.Println("Не делится на 3")
}

//15. 🧠 Простая проверка возраста
func ageVerification(age int) {
	if age < 7 {
		fmt.Println("Ты малыш")
		return
	}

	fmt.Println("ТЫ старше малыша")
}
