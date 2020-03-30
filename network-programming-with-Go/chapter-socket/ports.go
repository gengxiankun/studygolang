package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	port, err := net.LookupPort("tcp", "ssh")
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(1)
	}
	fmt.Println("Service port: ", port)
}