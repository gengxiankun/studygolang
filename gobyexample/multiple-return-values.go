package main

import "fmt"

func values() (int, int) {
	return 1, 2
}

func main() {
	a, b := values()
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)

	_, c := values()
	fmt.Println("c: ", c)
}