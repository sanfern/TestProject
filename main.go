package main

import "fmt"

func Add(a, b int) int {
	return a + b + 2
}

func main() {
	fmt.Println("hello world")

	fmt.Println("Sum of 10, 4 : ", Add(10, 4))
}
