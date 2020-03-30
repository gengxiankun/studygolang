// 一个时间(Daytime)服务器
// 最简单的服务，我们可以建立是时间(Daytime)服务。
// 这是一个标准的互联网服务, 由RFC 867定义, 默认的端口13, 协议是TCP和UDP。
// 很遗憾, 对安全的偏执，几乎没有任何站点运行着时间(Daytime)服务器。
// 不过没关系，我们可以建立我们自己的。 (对于那些有兴趣, 你可以在你的系统安装inetd, 你通常可以得到一个时间(Daytime)服务器。)

// func ListenTCP(net string, laddr *TCPAddr) (l *TCPListener, err os.Error)
// func (l *TCPListener) Accept() (c Conn, err os.Error)
// net参数可以设置为字符串"tcp", "tcp4"或者"tcp6"中的一个。
package main

import (
	"net"
	"os"
	"fmt"
	"time"
)

func main() {
	// 如果你想监听所有网络接口，IP地址应设置为0，或如果你只是想监听一个简单网络接口，IP地址可以设置为该网络的地址。
	// 如果端口设置为0，O/S会为你选择一个端口。否则，你可以选择你自己的。
	// 需要注意的是，在Unix系统上，除非你是监控系统，否则不能监听低于1024的端口，小于128的端口是由IETF标准化。
	// 该示例程序选择端口1200没有特别的原因。TCP地址如下":1200" - 所有网络接口, 端口1200。
	service := ":1200"
	addr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	// 在一个服务器上注册并监听一个端口。
	listener, err := net.ListenTCP("tcp", addr)
	checkError(err)

	for {
		// 然后它阻塞在一个"accept"操作，并等待客户端连接。
		// 当一个客户端连接, accept调用返回一个连接(connection)对象。
		conn, err := listener.Accept()
		checkError(err)

		// 时间(Daytime)服务非常简单，只是将当前时间写入到客户端, 关闭该连接，并继续等待下一个客户端。
		dayTime := time.Now().String()
		conn.Write([]byte(dayTime))
		conn.Close()
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(1)
	}
}