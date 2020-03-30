// 常规的通过通道发送和接收数据是阻塞的。然而，我们可以
// 使用带一个 `default` 子句的 `select` 来实现_非阻塞_ 的
// 发送、接收，甚至是非阻塞的多路 `select`。

package main

import "fmt"

func main() {
    messages := make(chan string)
    // signals := make(chan bool)

    // 这里是一个非阻塞接收的例子。如果在 `messages` 中s
    // 存在，然后 `select` 将这个值带入 `<-messages` `case`
    // 中。如果不是，就直接到 `default` 分支中。
    select {
    case msg := <-messages:
        fmt.Println("received message", msg)
    default:
        fmt.Println("no message received")
    }

    // 一个非阻塞发送的实现方法和上面一样。
    msg := "hi"
    select {
    case messages <- msg:
        fmt.Println("sent message", msg)
    default:
        fmt.Println("no message sent")
    }
}