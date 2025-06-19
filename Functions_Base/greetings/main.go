package main

import "fmt"

func main() {
	var name string
	fmt.Print("Введите имя: ")
	fmt.Scan(&name)
	fmt.Println(greetings(name))

}

func greetings(name string) string {
	return "Привет, " + name + "!"
}
