package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"time"
)

var (
	result []int      = make([]int, 0) // 素数切片
	lock   sync.Mutex                  // 全局锁
)

func exits(slice []int, n int) bool {
	for _, value := range slice {
		if value == n {
			return true
		}
	}

	return false
}

func is_prime(n int) {
	is_prime := true
	for i := 2; i < int(math.Sqrt(float64(n)))+1; i++ {
		if n%i == 0 {
			is_prime = false
			break
		}
	}

	if is_prime && !exits(result, n) {
		lock.Lock() // 请求锁
		result = append(result, n)
		lock.Unlock() // 释放锁
	}
}

func main() {
	num := runtime.NumCPU()
	runtime.GOMAXPROCS(num - 1)
	for i := 2; i < 2000; i++ {
		// 开启将近2000个协程，判断素数
		go is_prime(i)
	}
	fmt.Println("value")
	time.Sleep(10 * time.Second) // 主线程等待10秒

	lock.Lock() // 遍历的时候依旧要利用锁进行同步控制
	for _, value := range result {
		fmt.Println(value)
	}
	lock.Unlock()
}
