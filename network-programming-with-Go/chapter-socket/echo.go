// "echo"是另一种简单的IETF服务。只是读取客户端数据，并将其发送回去:
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	service := ":1201"
	addr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", addr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		// 客户端很容易使用go-routine扩展，实现多线程服务器 Multi-threaded server
		go handelClient(conn)		
	}
}

func handelClient(conn net.Conn) {
	var buf [512]byte

	// 读取512字节数据
	n, err := conn.Read(buf[0:])
	if err != nil {
		return
	}

	// 写入读取的n个字节
	_, err2 := conn.Write(buf[0:n])
	if err2 != nil {
		return
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatel error: ", err.Error())
		os.Exit(1)
	}
}