package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for {
		nums := rand.Intn(1_000_000)
		fmt.Printf("S%06d\n", nums)
		time.Sleep(500 * time.Millisecond)
	}
}
