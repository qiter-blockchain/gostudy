package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	option := redis.DialPassword("123456")
	var pool *redis.Pool // 全局连接池指针
	pool = &redis.Pool{
		MaxIdle:     8,   // 最大空闲连接数
		MaxActive:   0,   // 最大连接数,0表示不限
		IdleTimeout: 100, // 最大空闲时间
		Dial: func() (redis.Conn, error) { // 产生连接的函数
			return redis.Dial("tcp", "192.168.0.26:6379", option)
		},
	}
	conn := pool.Get() // 获取连接
	var err error
	defer pool.Close()                                    // 连接池关闭
	_, err = conn.Do("HSet", "userhash01", "name", "szc") // 操作、哈希名、键、值
	_, err = conn.Do("HSet", "userhash01", "age", 23)

	r, err := redis.String(conn.Do("HGet", "userhash01", "name")) // 从哈希userhash01中读取name
	fmt.Println("hash name = ", r)
	age, err := redis.Int(conn.Do("HGet", "userhash01", "age")) // 读取int
	fmt.Println("hash age = ", age)
	fmt.Println(err)
}
