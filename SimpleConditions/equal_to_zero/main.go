package main

import "fmt"

func main() {
	equalToZero(1)
}

//11. 🔍 Равно ли нулю
func equalToZero(number int) {
	if number == 0 {
		fmt.Println("Число равно нулю")
		return
	}

	fmt.Println("Число не равно нулю")
}
