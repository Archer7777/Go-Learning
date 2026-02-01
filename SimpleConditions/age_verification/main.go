package main

import "fmt"

func main() {
	ageVerification(7)
}

//15. 🧠 Простая проверка возраста
func ageVerification(age int) {
	if age < 7 {
		fmt.Println("Ты малыш")
		return
	}

	fmt.Println("ТЫ старше малыша")
}
