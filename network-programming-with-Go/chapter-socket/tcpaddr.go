// TCPAddr类型包含一个IP和一个port的结构:
// type TCPAddr struct {
//     IP   IP
//     Port int
// }
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// 函数ResolveTCPAddr用来创建一个TCPAddr
	addr, err := net.ResolveTCPAddr("tcp", "www.baidu.com:80")
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(1)
	}
	fmt.Println(addr)
}
