package main

import "fmt"

func add(a int,  b int) int {

	return a + b
}

func main() {

	les := add(1, 2)
	fmt.Println("1 + 2 = ", les)
}