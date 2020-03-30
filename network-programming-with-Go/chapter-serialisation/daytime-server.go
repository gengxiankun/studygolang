// 将当前时间作为ASN.Time类型时间来传送的TCP服务器
package main

import (
	"fmt"
	"os"
	"net"
	"encoding/asn1"
	"time"
)

func main() {
	service := ":1200"
	addr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", addr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		daytime := time.Now()
		mdata, err := asn1.Marshal(daytime)
		conn.Write(mdata)
		conn.Close()
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatel error: ", err.Error())
		os.Exit(1)
	}
}