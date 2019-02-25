package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println("Now i have", add(42, 13), "years old")
}
