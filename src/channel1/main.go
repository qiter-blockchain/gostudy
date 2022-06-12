package main

//原始数据管道、素数结果管道和协程状态管道
var (
	int_chan1 chan<- int = make(chan int, 8) // 只写
	int_chan2 <-chan int = make(chan int, 8) // 只读
)

func main() {

}
