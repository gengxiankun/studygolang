// 主机查询
package main

import (
	"net"
	"fmt"
	"os"
)

func main() {
	// ResolveIPAddr函数将对某个主机名执行DNS查询，并返回一个简单的IP地址。
	// 然而，通常主机如果有多个网卡，则可以有多个IP地址。它们也可能有多个主机名，作为别名。
	// 这些地址将会被归类为“canonical”主机名。
	// 如果你想找到的规范名称，使用func LookupCNAME(name string) (cname string, err os.Error)
	addrs, err := net.LookupHost("www.baidu.com")
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(1)
	}

	for _, s := range addrs {
		fmt.Println(s)
	}
}