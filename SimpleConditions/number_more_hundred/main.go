package main

import "fmt"

func main() {
	numberMoreHundred(120)
}

//12. 🪞 Число больше 100?
func numberMoreHundred(number int) {
	if number > 100 {
		fmt.Println("Много")
	} else {
		fmt.Println("Норм")
	}
}
