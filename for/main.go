package main

import "fmt"

func main() {
	sum := 0

	for index := 0; index < 10; index++ {
		sum += 1
	}

	fmt.Println(sum)
}