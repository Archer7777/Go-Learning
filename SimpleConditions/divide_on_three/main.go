package main

import "fmt"

func main() {
	divideOnThree(10)
}

//14. 📛 Делится на 3
func divideOnThree(number int) {
	if number%3 == 0 {
		fmt.Println("Делится")
		return
	}

	fmt.Println("Не делится на 3")
}
