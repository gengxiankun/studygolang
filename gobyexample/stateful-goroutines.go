package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type readOp struct {
	key int
	resp chan int
}

type writeOp struct {
	key int
	val int
	resp chan bool
}

func main() {

	// 和前面一样，我们将计算我们执行操作的次数。
	var ops int64

	// reads 和 writes 通道分别将被其他 Go 协程用来发布读和写请求。
	reads := make(chan *readOp)
	writes := make(chan *writeOp)

	// 这个就是拥有 state 的那个 Go 协程，
	// 和前面例子中的map一样，不过这里是被这个状态协程私有的。
	// 这个 Go 协程反复响应到达的请求。
	// 先响应到达的请求，然后返回一个值到响应通道 resp 来表示操作成功（或者是 reads 中请求的值）
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// 启动 100 个 Go 协程通过 reads 通道发起对 state 所有者Go 协程的读取请求。
	// 每个读取请求需要构造一个 readOp，发送它到 reads 通道中，
	// 并通过给定的 resp 通道接收结果。
	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := &readOp{
					key: rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddInt64(&ops, 1)
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := &writeOp{
					key: rand.Intn(5),
					val: rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddInt64(&ops, 1)
			}
		}()
	}

	time.Sleep(time.Second)

	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops:", opsFinal)
}