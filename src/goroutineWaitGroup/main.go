package main

import (
	"fmt"
	"runtime"
	"sync"
)

func Show(i int) {
	fmt.Println("show: ", i)
	defer wp.Done() //执行完任务协程-1
}

var wp sync.WaitGroup //使用WaitGroup 等待各个协程执行完

func main() {
	num := runtime.NumCPU() // 获取cpu核数
	fmt.Println("CPU count:", num)
	runtime.GOMAXPROCS(num - 1) // 设置最大并发数
	for i := 0; i < 10; i++ {
		go Show(i) // 开启10个 go协程执行任务
		wp.Add(1)  //协程+1
	}
	wp.Wait() //阻塞 直到协程组内协程数为0时往下执行
	fmt.Println("show Done")

}
