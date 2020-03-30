package main

import "fmt"

func zeroval(ival int) {
	ival = 0 
}

func zeropter(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("inital: ", i)

	zeroval(i)
	fmt.Println("zeroval: ", i)

	zeropter(&i)
	fmt.Println("zeropter: ", i)

	fmt.Println("pointer: ", &i)
}