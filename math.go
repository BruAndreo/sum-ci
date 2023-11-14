package main

import "fmt"

func main() {
	fmt.Println(sum(10, 10))
}

func sum(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func times(a, b int) int {
	return a * b
}

func sumX(a, b int) int {
	return a + b + a
}

func sumX2(a, b int) int {
	return a + b + a + b
}

func sumX3(a, b int) int {
	return a + b + a + b + a
}
