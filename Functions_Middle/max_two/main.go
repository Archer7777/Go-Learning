package main

func main() {
	max(5, 100)
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
