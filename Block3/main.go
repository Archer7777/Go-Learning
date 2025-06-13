package main

import "fmt"

func main() {
	equalToZero(1)
	numberMoreHundred(120)
	comparingTwoNumbers(10, 9)
	divideOnThree(10)
	ageVerification(7)
}

func equalToZero(number int) {
	if number == 0 {
		fmt.Println("Число равно нулю")
		return
	}

	fmt.Println("Число не равно нулю")
}

func numberMoreHundred(number int) {
	if number > 100 {
		fmt.Println("Много")
	} else {
		fmt.Println("Норм")
	}
}

func comparingTwoNumbers(a, b int) {
	if b > a {
		fmt.Println("Второе больше")
		return
	}

	fmt.Println("Первое больше")
}

func divideOnThree(number int) {
	if number%3 == 0 {
		fmt.Println("Делится")
		return
	}

	fmt.Println("Не делится на 3")
}

func ageVerification(age int) {
	if age < 7 {
		fmt.Println("Ты малыш")
		return
	}

	fmt.Println("ТЫ старше малыша")
}
