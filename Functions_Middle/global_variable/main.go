package main

import "fmt"

var counter int

func main() {
	increment()
	increment()
	increment()
	increment()

}

func increment() int {
	counter++
	fmt.Println(counter)
	return counter
}
