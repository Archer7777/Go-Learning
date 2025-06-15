package main

import (
	"fmt"
)

func main() {
	celciusInFarenheit(32)

}

// 18. 🌡️ Температура в Фаренгейтах
func celciusInFarenheit(celcius float64) {
	fmt.Println(celcius*1.8 + 32)
}
