package main

import (
	"net"
	"fmt"
	"os"
)

func main() {
	service := "www.baidu.com:80"
	addr, err := net.ResolveUDPAddr("udp", service)
	checkError(err)

	conn, err := net.DialUDP("udp", nil, addr)
	checkError(err)

	_, err = conn.Write([]byte("anything"))
	checkError(err)

	var buf [512]byte
	n, err := conn.Read(buf[0:])
	checkError(err)

	fmt.Println(buf[0:n])

	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatel error: ", err.Error())
		os.Exit(1)
	}
}