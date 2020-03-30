package main

import "fmt"
import "time"

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {
	done := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		worker(done)		
	}

	<-done
}