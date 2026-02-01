package main

func main() {
	getInHand(100000)

}

// 19. 💵 Сколько получишь на руки
func getInHand(salary float64) float64 {
	return salary - salary*0.13
}
