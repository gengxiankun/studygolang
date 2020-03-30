package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is old")
	}

	if 8%4 == 0 {
		fmt.Println("8%4 = ", 8%4)
	}

	if num := 9; num < 0 {
		fmt.Println("num < 0")
	} else if num < 10 {
		fmt.Println("num < 10")
	}
}