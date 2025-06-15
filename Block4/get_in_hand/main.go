package main

import (
	"fmt"
)

func main() {
	getInHand(100000)

}

// 19. 💵 Сколько получишь на руки
func getInHand(salary float64) {
	fmt.Println(salary - salary*0.13)
}
