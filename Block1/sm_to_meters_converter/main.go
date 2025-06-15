package main

import "fmt"

func main() {
	smToMeters(250)
}

//3. 📏 Перевод сантиметров в метры
func smToMeters(sm float64) {
	fmt.Printf("Вывод: %g м\n", (sm / 100))
}
