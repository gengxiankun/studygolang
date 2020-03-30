package main

import "fmt"

func sum(nums ...int) {
	fmt.Println("len: ", len(nums))
	var total int

	for _, num := range nums {
		total = total + num
	}
	fmt.Println("total: ", total)
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(nums)

	sum(nums...)
	sum(1, 2)
}