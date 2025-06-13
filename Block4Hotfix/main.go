package main

import (
	"fmt"
)

func main() {
	// evenOrNotEven()
	// adult()
	// seasons()
	// maxTwoNumbers()
	// positiveNegativeOrZero()
	grade()
}

// // 🧠 1. Четное или нечетное
// func evenOrNotEven() {
// 	var number int
// 	fmt.Scanf("%d", &number)
// 	if number%2 == 0 {
// 		fmt.Println("Число четное")
// 	} else {
// 		fmt.Println("Число нечетное")
// 	}
// }

// // // // // 🔞 2. Совершеннолетие
// func adult() {
// 	var age int
// 	fmt.Scanf("%d", &age)
// 	if age >= 18 {
// 		fmt.Println("Доступ разрешён")
// 		return
// 	}

// 	fmt.Println("Доступ запрещён")
// }

// // // // 📅 3. Определение сезона
// func seasons() {
// 	var monthNumber int
// 	fmt.Scanf("%d", &monthNumber)
// 	switch monthNumber {
// 	case 12, 1, 2:
// 		fmt.Println("Зима")
// 	case 3, 4, 5:
// 		fmt.Println("Весна")
// 	case 6, 7, 8:
// 		fmt.Println("Лето")
// 	case 9, 10, 11:
// 		fmt.Println("Осень")
// 	}
// }

// // // 🧮 4. Максимум из двух чисел
// func maxTwoNumbers() {
// 	var number1, number2 int
// 	fmt.Scanf("%d %d", &number1, &number2)
// 	if number1 > number2 {
// 		fmt.Println(number1)
// 	} else if number1 == number2 {
// 		fmt.Println("Числа равны")
// 	} else {
// 		fmt.Println(number2)
// 	}
// }

// // 🔢 5. Число — положительное, отрицательное или ноль
// func positiveNegativeOrZero() {
// 	var number int
// 	fmt.Scanf("%d", &number)
// 	if number < 0 {
// 		fmt.Println("Число отрицательное")
// 	} else if number > 0 {
// 		fmt.Println("Число положительное")
// 	} else {
// 		fmt.Println("Число Ноль")
// 	}
// }

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
