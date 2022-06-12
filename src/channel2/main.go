package main

import "fmt"

var (
	write_chan chan int  = make(chan int, 50) // 数据管道，整型管道，容量50(不可扩容)
	exit_chan  chan bool = make(chan bool, 1) // 状态管道，布尔型管道，容量1(不可扩容)
)

func write_data() {
	for i := 0; i < 50; i++ {
		write_chan <- i // 往管道里写数据
		fmt.Println("write data: ", i)
	}

	close(write_chan)
	// 关闭管道不影响读，只影响写
}

func read_data() {
	for {
		v, ok := <-write_chan // 从管道里读数据，返回具体数据和成功与否。如果管道为空，就会阻塞
		if !ok {              // 如果管道为空，则ok为false
			break
		}
		fmt.Println("read data: ", v)
	}

	exit_chan <- true
	close(exit_chan)
}

//主线程负责开启两个协程，并监视状态管道
func main() {
	go write_data()
	go read_data()

	for {
		_, ok := <-exit_chan
		if !ok {
			break
		}
	}

}
