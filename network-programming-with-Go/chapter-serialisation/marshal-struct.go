// 编、解组结构类型
package main

import (
	"fmt"
	"encoding/asn1"
)

type T struct {
	Field1 int
	Field2 int
}

func main() {
	// 使用变量
	t1 := T{Field1: 1, Field2: 2}
	mdata, _ := asn1.Marshal(t1)
	fmt.Println("Marshal: ", mdata)

	var newt1 T
	asn1.Unmarshal(mdata, &newt1)
	fmt.Println("Unmarshal: ", newt1)

	// 使用指针
	// 恰当地的使用指针与变量能让代码工作得更好。
	var t2 = new(T)
	t2.Field1 = 1
	t2.Field2 = 2
	mdata, _ = asn1.Marshal(*t2)
	fmt.Println("Marshal: ", mdata)

	var newt2 = new(T)
	asn1.Unmarshal(mdata, newt2)
	fmt.Println("Unmarshal: ", newt2)
}