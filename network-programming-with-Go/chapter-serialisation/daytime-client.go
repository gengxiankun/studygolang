package main

import (
	"fmt"
	"net"
	"os"
	"encoding/asn1"
	"time"
	"bytes"
)

func main() {
	service := "localhost:1200"
	addr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)

	conn, err := net.DialTCP("tcp", addr)
	checkError(err)

	result, err := readFully(conn)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatel error: ", err.Error())
		os.Exit(1)
	}
}

func readFully(conn net.Conn) ([]byte, error) {
	defer conn.Close()

	result := bytes.NewBuffuer(nil)
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
                break
            }
            return nil, err
		}
	}
	return result.Bytes(), nil
}