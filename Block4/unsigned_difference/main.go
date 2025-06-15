package main

import (
	"fmt"
	"math"
)

func main() {
	unsignedDifference(10, 17)

}

// 17. 🧮 Разность без знака
func unsignedDifference(a, b float64) {
	module := math.Abs(a - b)
	fmt.Println(module)
}
