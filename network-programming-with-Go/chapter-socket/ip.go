// "net"包定义了许多类型, 函数，方法用于Go网络编程。
// IP类型被定义为一个字节数组：type IP []byte
// 有几个函数来处理一个IP类型的变量, 但是在实践中你很可能只用到其中的一些。
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// ParseIP(String)函数将获取逗号分隔的IPv4或者冒号分隔的IPv6地址
	addr := net.ParseIP("127.0.0.1")
	if addr == nil {
		fmt.Println("Invalid address")
		os.Exit(1)
	}
	// ParseIP(String)函数将获取逗号分隔的IPv4或者冒号分隔的IPv6地址
	// 您可能无法取回你期望的: 字符串 0:0:0:0:0:0:0:1是::1
	fmt.Println("The address is ", addr.String())
}