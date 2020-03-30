package main

import "fmt"

type people struct {
	name string
	age  int
}

func main() {
	fmt.Println(people{"Bob", 20})
	fmt.Println(people{name: "Alice", age: 21})

	fmt.Println(people{name: "Fred"})

	fmt.Println(&people{name: "Ann", age: 40})

	s := people{name: "Sean", age: 50}
	fmt.Println(s)

	sp := &s
	fmt.Println(sp)

	sp.age = 51
	fmt.Println(s)
}