// 为了处理掩码操作，有下面类型：type IPMask []byte
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	addr := net.ParseIP("127.0.0.1")
	if addr == nil {
		fmt.Println("Invalid address")
		os.Exit(1)
	}

	// 这是一个IP的方法返回默认的掩码
	mask := addr.DefaultMask()
	// 一个掩码可以使用一个IP地址的方法，找到该IP地址的网络
	network := addr.Mask(mask)
	ones, bits := mask.Size()

	fmt.Println("Address is ", addr.String(),
        " Default mask length is ", bits,
        "Leading ones count is ", ones,
        "Mask is (hex) ", mask.String(),
        " Network is ", network.String())
}