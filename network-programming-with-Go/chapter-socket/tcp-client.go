// net.TCPConn是允许在客户端和服务器之间的全双工通信的Go类型。两种主要方法是：
//  - func (c *TCPConn) Write(b []byte) (n int, err os.Error)
//  - func (c *TCPConn) Read(b []byte) (n int, err os.Error)   
// TCPConn被客户端和服务器用来读写消息。
package main

import (
	"fmt"
	"net"
	"os"
	"io/ioutil"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", "www.baidu.com:80")
	checkError(err)

	// 一旦客户端已经建立TCP服务, 就可以和对方设备"通话"了. 
	// 如果成功，该调用返回一个用于通信的TCPConn。客户端和服务器通过它交换消息。
	// 通常情况下，客户端使用TCPConn写入请求到服务器, 并从TCPConn的读取响应。
	// 持续如此，直到任一（或两者）的两侧关闭连接。客户端使用该函数建立一个TCP连接。

	// func DialTCP(net string, laddr, raddr *TCPAddr) (c *TCPConn, err os.Error)
	//  - 其中laddr是本地地址，通常设置为nil
	//  - raddr是一个服务的远程地址, net是一个字符串，根据您是否希望是一个TCPv4连接，TCPv6连接来设置为"tcp4", "tcp6"或"tcp"中的一个，当然你也可以不关心链接形式。
	conn, err := net.DialTCP("tcp", nil, addr)
	checkError(err)

	// 客户端可能发送的消息之一就是“HEAD”消息。
	// 这用来查询服务器的信息和文档信息。 
	// 服务器返回的信息，不返回文档本身。发送到服务器的请求可能是
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)

	// 值得一提的是,如何从服务端读取数据。
	// 在这种情况下，读本质上是一个单一的来自服务器的响应，这将终止文件结束的连接。
	// 但是，它可能包括多个TCP数据包，所以我们需要不断地读，直到文件的末尾。
	// 在io/ioutil下的ReadAll函数考虑了这些问题，并返回完整响应。(感谢Roger Peppe在golang-nuts上的邮件列表。)。
	result, err := ioutil.ReadAll(conn)
	checkError(err)

	fmt.Println(string(result))
}

// 第一个要注意的点是近乎多余的错误检查。
// 因为正常情况下，网络程序失败的机会大大超过单机的程序。在客户端，服务器端或任何路由和中间交换上，硬件可能失败；通信可能会被防火墙阻塞;因网络负载可能会出现超时;当客户端联系服务器，服务器可能会崩溃，下列检查是必须的：
//  - 指定的地址中可能存在语法错误
//  - 尝试连接到远程服务可能会失败。例如, 所请求的服务可能没有运行, 或者有可能是主机没有连接到网络
//  - 虽然连接已经建立，如果连接突然丢失也可能导致写失败，或网络超时
//  - 同样，读操作也可能会失败
func checkError(err error) {
	if err != nil {
		fmt.Println("Fatel error: ", err.Error())
		os.Exit(1)
	}
}

