// Go 支持在结构体类型中定义方法 。
package main

import "fmt"

type rect struct {
	width, height int
}

// 这里的 area 方法有一个接收器类型 rect。
func (r *rect) area() int {
	return r.width * r.height
}

func main() {
	r := rect{width: 20, height: 10}
	fmt.Println("area: ", r.area())

	rp := &r
	fmt.Println("area: ", rp.area())
}