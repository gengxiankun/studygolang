// 在net包的许多函数和方法会返回一个指向IPAddr的指针。这不过只是一个包含IP类型的结构体。
// type IPAddr {
//     IP IP
// }
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// 这种类型的主要用途是通过IP主机名执行DNS查找。
	addr, err := net.ResolveIPAddr("ip", "www.baidu.com")
	if err != nil {
		fmt.Println("Resolution error ", err.Error())
		os.Exit(1)
	}

	fmt.Println("Resolved address is ", addr.String())
}