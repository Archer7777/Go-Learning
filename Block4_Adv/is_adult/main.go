package main

import (
	"fmt"
)

func main() {
	adult()
}

// 🔞 2. Совершеннолетие
func adult() {
	var age int
	fmt.Scanf("%d", &age)
	if age >= 18 {
		fmt.Println("Доступ разрешён")
		return
	}

	fmt.Println("Доступ запрещён")
}
