package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i = i + 1
		return i
	}
}

func main() {

	nestInt := intSeq()

	fmt.Println(nestInt())
	fmt.Println(nestInt())
	fmt.Println(nestInt())
}