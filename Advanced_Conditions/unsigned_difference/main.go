package main

import (
	"math"
)

func main() {
	unsignedDifference(10, 17)

}

// 17. 🧮 Разность без знака
func unsignedDifference(a, b float64) float64 {
	module := math.Abs(a - b)
	return module
}
